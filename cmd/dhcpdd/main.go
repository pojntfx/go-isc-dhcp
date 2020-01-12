package main

import (
	"fmt"
	"github.com/pojntfx/godhcpd/pkg/svc"
	"github.com/pojntfx/godhcpd/pkg/workers"
	uuid "github.com/satori/go.uuid"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	id := uuid.NewV4().String()
	stateDir := filepath.Join(os.TempDir(), "godhcpd", "dhcpd", id)
	binaryDir := filepath.Join(os.TempDir(), "dhcpd")

	dhcpdManager := svc.DHCPDManager{
		BinaryDir: binaryDir,
	}
	if err := dhcpdManager.Extract(); err != nil {
		fmt.Println(err)
	}

	dhcpd := workers.DHCPD{
		Subnets: []workers.Subnet{
			{
				Network: "192.168.1.0",
				Netmask: "255.255.255.0",
				Range: workers.Range{
					Start: "192.168.1.10",
					End:   "192.168.1.100",
				},
			},
		},
		BinaryDir: dhcpdManager.BinaryDir,
		ID:        id,
		StateDir:  stateDir,
		Device:    "edge0",
	}

	if err := dhcpd.Configure(); err != nil {
		fmt.Println(err)
	}

	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interrupt

		// Allow manually killing the process
		go func() {
			<-interrupt

			os.Exit(1)
		}()

		if err := dhcpd.Stop(); err != nil {
			fmt.Println(err)
		}
	}()

	if err := dhcpd.Start(); err != nil {
		fmt.Println(err)
	}

	if err := dhcpd.Wait(); err != nil {
		fmt.Println(err)
	}
}
