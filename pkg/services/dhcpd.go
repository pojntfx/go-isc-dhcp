package services

//go:generate sh -c "mkdir -p ../api/proto/v1 && protoc --go_out=paths=source_relative,plugins=grpc:../api/proto/v1 -I=../../api/proto/v1 ../../api/proto/v1/*.proto"

import (
	"context"
	"os"
	"path/filepath"

	api "github.com/pojntfx/go-isc-dhcp/pkg/api/proto/v1"
	"github.com/pojntfx/go-isc-dhcp/pkg/workers"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DHCPDManager manages dhcp servers.
type DHCPDManager struct {
	api.UnimplementedDHCPDManagerServer
	BinaryDir     string
	StateDir      string
	DHCPDsManaged map[string]*workers.DHCPD
}

func (m *DHCPDManager) getReplyDHCPDManagerFromDHCPDManaged(id string, DHCPD *workers.DHCPD) *api.DHCPDManaged {
	var subnetsForReply []*api.Subnet
	for _, subnet := range DHCPD.Subnets {
		subnetForReply := &api.Subnet{
			Network:    subnet.Network,
			Netmask:    subnet.Netmask,
			NextServer: subnet.NextServer,
			Filename:   subnet.Filename,
			Range: &api.Range{
				Start: subnet.Range.Start,
				End:   subnet.Range.End,
			},
		}

		subnetsForReply = append(subnetsForReply, subnetForReply)
	}

	return &api.DHCPDManaged{
		Id:      id,
		Device:  DHCPD.Device,
		Subnets: subnetsForReply,
	}
}

// Create creates a dhcp server.
func (m *DHCPDManager) Create(_ context.Context, args *api.DHCPD) (*api.DHCPDManagedId, error) {
	id := uuid.NewV4().String()

	subnets := args.GetSubnets()
	var subnetsForWorker []workers.Subnet
	for _, subnet := range subnets {
		subnetsForWorker = append(subnetsForWorker, workers.Subnet{
			Network:           subnet.GetNetwork(),
			Netmask:           subnet.GetNetmask(),
			NextServer:        subnet.GetNextServer(),
			Filename:          subnet.GetFilename(),
			Routers:           subnet.GetRouters(),
			DomainNameServers: subnet.GetDomainNameServers(),
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

	return &api.DHCPDManagedId{
		Id: id,
	}, nil
}

// List lists the managed dhcp servers.
func (m *DHCPDManager) List(_ context.Context, args *api.DHCPDManagerListArgs) (*api.DHCPDManagerListReply, error) {
	log.Info("Listing dhcp servers")

	var DHCPDsManaged []*api.DHCPDManaged
	for id, DHCPD := range m.DHCPDsManaged {
		DHCPDsManaged = append(DHCPDsManaged, m.getReplyDHCPDManagerFromDHCPDManaged(id, DHCPD))
	}

	return &api.DHCPDManagerListReply{
		DHCPDsManaged: DHCPDsManaged,
	}, nil
}

// Get gets one of the managed dhcp servers.
func (m *DHCPDManager) Get(_ context.Context, args *api.DHCPDManagedId) (*api.DHCPDManaged, error) {
	log.Info("Getting dhcp server")

	var DHCPDManaged *api.DHCPDManaged

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
func (m *DHCPDManager) Delete(_ context.Context, args *api.DHCPDManagedId) (*api.DHCPDManagedId, error) {
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

	return &api.DHCPDManagedId{
		Id: id,
	}, nil
}

// Extract extracts the ISC DHCP server binary.
func (m *DHCPDManager) Extract() error {
	binaryFile, err := os.Create(m.BinaryDir)
	if err != nil {
		return err
	}

	if _, err = binaryFile.Write(workers.EmbeddedDHCPD); err != nil {
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
