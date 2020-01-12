package main

import (
	"fmt"
	godhcpd "github.com/pojntfx/godhcpd/pkg/proto/generated"
	"github.com/pojntfx/godhcpd/pkg/svc"
	"github.com/pojntfx/godhcpd/pkg/workers"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	binaryDir := filepath.Join(os.TempDir(), "dhcpd")

	listener, err := net.Listen("tcp", "localhost:30001")
	if err != nil {
		fmt.Println(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	DHCPDService := svc.DHCPDManager{
		BinaryDir:     binaryDir,
		StateDir:      filepath.Join(os.TempDir(), "godhcpd", "dhcpd"),
		DHCPDsManaged: make(map[string]*workers.DHCPD),
	}

	godhcpd.RegisterDHCPDManagerServer(server, &DHCPDService)

	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interrupt

		// Allow manually killing the process
		go func() {
			<-interrupt

			os.Exit(1)
		}()

		log.Info("Gracefully stopping server (this might take a few seconds)")

		msg := "Could not stop dhcp server"

		for _, DHCPD := range DHCPDService.DHCPDsManaged {
			if err := DHCPD.Stop(); err != nil {
				log.Fatal(msg, rz.Err(err))
			}
		}

		for _, DHCPD := range DHCPDService.DHCPDsManaged {
			if err := DHCPD.Wait(); err != nil {
				log.Fatal(msg, rz.Err(err))
			}
		}

		server.GracefulStop()
	}()

	if err := DHCPDService.Extract(); err != nil {
		fmt.Println(err)
	}

	log.Info("Starting server")

	if err := server.Serve(listener); err != nil {
		fmt.Println(err)
	}
}
