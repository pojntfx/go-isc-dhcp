package main

import (
	"fmt"
	"github.com/pojntfx/godhcpd/pkg/workers"
	uuid "github.com/satori/go.uuid"
	"os"
	"path/filepath"
)

func main() {
	id := uuid.NewV4().String()
	stateDir := filepath.Join(os.TempDir(), "godhcpd", "dhcpd", id)

	dhcpd := workers.DHCPD{
		ID:       id,
		StateDir: stateDir,
		Device:   "edge0",
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
	}

	if err := dhcpd.Configure(); err != nil {
		fmt.Println(err)
	}
}
