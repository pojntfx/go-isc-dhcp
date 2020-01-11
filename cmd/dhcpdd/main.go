package main

import (
	"github.com/pojntfx/godhcpd/pkg/workers"
	uuid "github.com/satori/go.uuid"
)

func main() {
	id := uuid.NewV4().String()

	dhcpd := workers.DHCPD{
		ID: id,
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

	dhcpd.Configure()
}
