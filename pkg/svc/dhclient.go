package svc

import (
	"context"
	godhcpd "github.com/pojntfx/godhcpd/pkg/proto/generated"
	"github.com/pojntfx/godhcpd/pkg/workers"
	"github.com/rakyll/statik/fs"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

// DHClientManager manages dhcp clients.
type DHClientManager struct {
	godhcpd.UnimplementedDHClientManagerServer
	BinaryDir        string
	DHClientsManaged map[string]*workers.DHClient
}

// Create creates a dhcp client.
func (m *DHClientManager) Create(_ context.Context, args *godhcpd.DHClient) (*godhcpd.DHClientManagedId, error) {
	id := uuid.NewV4().String()

	dhcpd := workers.DHClient{
		Device:    args.GetDevice(),
		BinaryDir: m.BinaryDir,
	}

	if err := dhcpd.Start(); err != nil {
		log.Error(err.Error())
	}

	go func(dhcpd *workers.DHClient) {
		log.Info("Starting dhcp client")

		_ = dhcpd.Wait()

		// Keep the dhcp client running
		for {
			if !dhcpd.IsScheduledForDeletion() {
				log.Info("Restarting dhcp client")

				if err := dhcpd.Start(); err != nil {
					log.Error(err.Error())
				}

				_ = dhcpd.Wait()
			} else {
				break
			}
		}
	}(&dhcpd)

	m.DHClientsManaged[id] = &dhcpd

	return &godhcpd.DHClientManagedId{
		Id: id,
	}, nil
}

// List lists the managed dhcp clients.
func (m *DHClientManager) List(_ context.Context, args *godhcpd.DHClientManagerListArgs) (*godhcpd.DHClientManagerListReply, error) {
	log.Info("Listing dhcp clients")

	var DHClients []*godhcpd.DHClientManaged
	for id, DHClient := range m.DHClientsManaged {
		DHClients = append(DHClients, &godhcpd.DHClientManaged{
			Id:     id,
			Device: DHClient.Device,
		})
	}

	return &godhcpd.DHClientManagerListReply{
		DHClientsManaged: DHClients,
	}, nil
}

// Get gets one of the managed dhcp clients.
func (m *DHClientManager) Get(_ context.Context, args *godhcpd.DHClientManagedId) (*godhcpd.DHClientManaged, error) {
	log.Info("Getting dhcp client")

	var DHClientManaged *godhcpd.DHClientManaged

	for id, DHClient := range m.DHClientsManaged {
		if id == args.GetId() {
			DHClientManaged = &godhcpd.DHClientManaged{
				Id:     id,
				Device: DHClient.Device,
			}
			break
		}
	}

	if DHClientManaged != nil {
		return DHClientManaged, nil
	}

	msg := "dhcp client not found"

	log.Error(msg)

	return nil, status.Errorf(codes.NotFound, msg)
}

// Delete deletes a dhcp client.
func (m *DHClientManager) Delete(_ context.Context, args *godhcpd.DHClientManagedId) (*godhcpd.DHClientManagedId, error) {
	id := args.GetId()

	DHClient := m.DHClientsManaged[id]
	if DHClient == nil {
		msg := "dhcp client not found"

		log.Error(msg)

		return nil, status.Errorf(codes.NotFound, msg)
	}

	log.Info("Stopping dhcp client")

	if err := DHClient.DisableAutoRestart(); err != nil { // Manually disable auto restart; disables crash recovery even if process is not running
		log.Error(err.Error())

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	if DHClient.IsRunning() {
		if err := DHClient.Stop(); err != nil { // Stop is sync, so no need to `.Wait()`
			log.Error(err.Error())

			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}

	delete(m.DHClientsManaged, id)

	return &godhcpd.DHClientManagedId{
		Id: id,
	}, nil
}

// Extract extracts the ISC DHCP client binary.
func (m *DHClientManager) Extract() error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}

	data, err := fs.ReadFile(statikFS, "/dhclient")
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

// Cleanup deletes the extracted ISC DHCP client binary.
func (m *DHClientManager) Cleanup() error {
	return os.Remove(m.BinaryDir)
}
