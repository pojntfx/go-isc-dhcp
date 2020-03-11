package dhcpd

import (
	"context"
	"os"
	"path/filepath"

	goISCDHCP "github.com/pojntfx/go-isc-dhcp/pkg/proto/generated"
	_ "github.com/pojntfx/go-isc-dhcp/pkg/svc/statikDhcpd" // Embedded ISC DHCP server binary
	"github.com/pojntfx/go-isc-dhcp/pkg/workers"
	"github.com/rakyll/statik/fs"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DHCPDManager manages dhcp servers.
type DHCPDManager struct {
	goISCDHCP.UnimplementedDHCPDManagerServer
	BinaryDir     string
	StateDir      string
	DHCPDsManaged map[string]*workers.DHCPD
}

func (m *DHCPDManager) getReplyDHCPDManagerFromDHCPDManaged(id string, DHCPD *workers.DHCPD) *goISCDHCP.DHCPDManaged {
	var subnetsForReply []*goISCDHCP.Subnet
	for _, subnet := range DHCPD.Subnets {
		subnetForReply := &goISCDHCP.Subnet{
			Network:    subnet.Network,
			Netmask:    subnet.Netmask,
			NextServer: subnet.NextServer,
			Filename:   subnet.Filename,
			Range: &goISCDHCP.Range{
				Start: subnet.Range.Start,
				End:   subnet.Range.End,
			},
		}

		subnetsForReply = append(subnetsForReply, subnetForReply)
	}

	return &goISCDHCP.DHCPDManaged{
		Id:      id,
		Device:  DHCPD.Device,
		Subnets: subnetsForReply,
	}
}

// Create creates a dhcp server.
func (m *DHCPDManager) Create(_ context.Context, args *goISCDHCP.DHCPD) (*goISCDHCP.DHCPDManagedId, error) {
	id := uuid.NewV4().String()

	subnets := args.GetSubnets()
	var subnetsForWorker []workers.Subnet
	for _, subnet := range subnets {
		subnetsForWorker = append(subnetsForWorker, workers.Subnet{
			Network:    subnet.GetNetwork(),
			Netmask:    subnet.GetNetmask(),
			NextServer: subnet.GetNextServer(),
			Filename:   subnet.GetFilename(),
			Range: workers.Range{
				Start: subnet.GetRange().GetStart(),
				End:   subnet.GetRange().GetEnd(),
			},
		})
	}

	dhcpd := workers.DHCPD{
		Subnets:   subnetsForWorker,
		Device:    args.GetDevice(),
		ID:        id,
		BinaryDir: m.BinaryDir,
		StateDir:  filepath.Join(m.StateDir, id),
	}

	if err := dhcpd.Configure(); err != nil {
		log.Error(err.Error())

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	if err := dhcpd.Start(); err != nil {
		log.Error(err.Error())
	}

	go func(dhcpd *workers.DHCPD) {
		log.Info("Starting dhcp server")

		_ = dhcpd.Wait()

		// Keep the dhcp server running
		for {
			if !dhcpd.IsScheduledForDeletion() {
				log.Info("Restarting dhcp server")

				if err := dhcpd.Start(); err != nil {
					log.Error(err.Error())
				}

				_ = dhcpd.Wait()
			} else {
				break
			}
		}
	}(&dhcpd)

	m.DHCPDsManaged[id] = &dhcpd

	return &goISCDHCP.DHCPDManagedId{
		Id: id,
	}, nil
}

// List lists the managed dhcp servers.
func (m *DHCPDManager) List(_ context.Context, args *goISCDHCP.DHCPDManagerListArgs) (*goISCDHCP.DHCPDManagerListReply, error) {
	log.Info("Listing dhcp servers")

	var DHCPDsManaged []*goISCDHCP.DHCPDManaged
	for id, DHCPD := range m.DHCPDsManaged {
		DHCPDsManaged = append(DHCPDsManaged, m.getReplyDHCPDManagerFromDHCPDManaged(id, DHCPD))
	}

	return &goISCDHCP.DHCPDManagerListReply{
		DHCPDsManaged: DHCPDsManaged,
	}, nil
}

// Get gets one of the managed dhcp servers.
func (m *DHCPDManager) Get(_ context.Context, args *goISCDHCP.DHCPDManagedId) (*goISCDHCP.DHCPDManaged, error) {
	log.Info("Getting dhcp server")

	var DHCPDManaged *goISCDHCP.DHCPDManaged

	for id, DHCPD := range m.DHCPDsManaged {
		if id == args.GetId() {
			DHCPDManaged = m.getReplyDHCPDManagerFromDHCPDManaged(id, DHCPD)
			break
		}
	}

	if DHCPDManaged != nil {
		return DHCPDManaged, nil
	}

	msg := "dhcp server not found"

	log.Error(msg)

	return nil, status.Errorf(codes.NotFound, msg)
}

// Delete deletes a dhcp server.
func (m *DHCPDManager) Delete(_ context.Context, args *goISCDHCP.DHCPDManagedId) (*goISCDHCP.DHCPDManagedId, error) {
	id := args.GetId()

	DHCPD := m.DHCPDsManaged[id]
	if DHCPD == nil {
		msg := "dhcp server not found"

		log.Error(msg)

		return nil, status.Errorf(codes.NotFound, msg)
	}

	log.Info("Stopping dhcp server")

	// Only stop; cleanup in interrupt handler
	if err := DHCPD.DisableAutoRestart(); err != nil { // Manually disable auto restart; disables crash recovery even if process is not running
		log.Error(err.Error())

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	if DHCPD.IsRunning() {
		if err := DHCPD.Stop(); err != nil { // Stop is sync, so no need to `.Wait()`
			log.Error(err.Error())

			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}

	delete(m.DHCPDsManaged, id)

	return &goISCDHCP.DHCPDManagedId{
		Id: id,
	}, nil
}

// Extract extracts the ISC DHCP server binary.
func (m *DHCPDManager) Extract() error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}

	data, err := fs.ReadFile(statikFS, "/dhcpd")
	if err != nil {
		return err
	}

	binaryFile, err := os.Create(m.BinaryDir)
	if err != nil {
		return err
	}

	if _, err = binaryFile.Write(data); err != nil {
		return err
	}
	defer binaryFile.Close()

	if err := os.Chmod(m.BinaryDir, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// Cleanup deletes the extracted ISC DHCP server binary.
func (m *DHCPDManager) Cleanup() error {
	return os.Remove(m.BinaryDir)
}
