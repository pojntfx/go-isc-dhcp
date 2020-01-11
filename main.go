/*
SPDX-License-Identifier: SPDX-3.0

This license is valid for all code in this repo.
*/

package main

import (
	"fmt"

	"github.com/pojntfx/godhcpd/pkg"
)

func main() {
	fmt.Println("Starting")

	dhcpServer := pkg.DHCPServer{
		ConfigurationFile: "/etc/dhcp/dhcpd.conf",
		Device:            "edge0",
	}

	if err := dhcpServer.Start(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Stopped")
}
