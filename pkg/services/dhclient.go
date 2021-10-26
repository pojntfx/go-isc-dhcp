package services

import (
	"context"
	"os"

	api "github.com/pojntfx/go-isc-dhcp/pkg/api/proto/v1"
	"github.com/pojntfx/go-isc-dhcp/pkg/workers"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DHClientManager manages dhcp clients.
type DHClientManager struct {
	api.UnimplementedDHClientManagerServer
	BinaryDir        string
	DHClientsManaged map[string]*workers.DHClient
}

// Create creates a dhcp client.
func (m *DHClientManager) Create(_ context.Context, args *api.DHClient) (*api.DHClientManagedId, error) {
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

	return &api.DHClientManagedId{
		Id: id,
	}, nil
}

// List lists the managed dhcp clients.
func (m *DHClientManager) List(_ context.Context, args *api.DHClientManagerListArgs) (*api.DHClientManagerListReply, error) {
	log.Info("Listing dhcp clients")

	var DHClients []*api.DHClientManaged
	for id, DHClient := range m.DHClientsManaged {
		DHClients = append(DHClients, &api.DHClientManaged{
			Id:     id,
			Device: DHClient.Device,
		})
	}

	return &api.DHClientManagerListReply{
		DHClientsManaged: DHClients,
	}, nil
}

// Get gets one of the managed dhcp clients.
func (m *DHClientManager) Get(_ context.Context, args *api.DHClientManagedId) (*api.DHClientManaged, error) {
	log.Info("Getting dhcp client")

	var DHClientManaged *api.DHClientManaged

	for id, DHClient := range m.DHClientsManaged {
		if id == args.GetId() {
			DHClientManaged = &api.DHClientManaged{
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
func (m *DHClientManager) Delete(_ context.Context, args *api.DHClientManagedId) (*api.DHClientManagedId, error) {
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

	return &api.DHClientManagedId{
		Id: id,
	}, nil
}

// Extract extracts the ISC DHCP client binary.
func (m *DHClientManager) Extract() error {
	binaryFile, err := os.Create(m.BinaryDir)
	if err != nil {
		return err
	}

	if _, err = binaryFile.Write(workers.EmbeddedDHClient); err != nil {
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
