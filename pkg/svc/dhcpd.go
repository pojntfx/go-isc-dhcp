package svc

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"
//go:generate sh -c "rm -rf dhcp statik; git clone https://gitlab.isc.org/isc-projects/dhcp.git; cd dhcp; ./configure; make; mkdir -p dist; cp server/dhcpd dist; cp client/dhclient dist; cd ../; statik -src dhcp/dist"

import (
	"context"
	godhcpd "github.com/pojntfx/godhcpd/pkg/proto/generated"
	_ "github.com/pojntfx/godhcpd/pkg/svc/statik" // Embedded ISC DHCP server binary
	"github.com/pojntfx/godhcpd/pkg/workers"
	"github.com/rakyll/statik/fs"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"path/filepath"
)

// DHCPDManager manages DHCP servers.
type DHCPDManager struct {
	godhcpd.UnimplementedDHCPDManagerServer
	BinaryDir     string
	StateDir      string
	DHCPDsManaged map[string]*workers.DHCPD
}

// Create creates a DHCP server.
func (m *DHCPDManager) Create(_ context.Context, args *godhcpd.DHCPDManagerCreateArgs) (*godhcpd.DHCPDManagerCreateReply, error) {
	id := uuid.NewV4().String()

	subnets := args.GetSubnets()
	var subnetsForWorker []workers.Subnet
	for _, subnet := range subnets {
		subnetsForWorker = append(subnetsForWorker, workers.Subnet{
			Network: subnet.GetNetwork(),
			Netmask: subnet.GetNetmask(),
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

		// Keep the dhcp server running
		for {
			_ = dhcpd.Wait()

			log.Info("Restarting dhcp server")

			if err := dhcpd.Start(); err != nil {
				log.Error(err.Error())
			}
		}
	}(&dhcpd)

	m.DHCPDsManaged[id] = &dhcpd

	return &godhcpd.DHCPDManagerCreateReply{
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
