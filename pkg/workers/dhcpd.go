package workers

import (
	"fmt"
	"os"
	"path/filepath"
)

// DHCPD is a DHCP server.
type DHCPD struct {
	Subnets       []Subnet
	ID            string
	StateDir      string
	Device        string
	configFileDir string
	leasesFileDir string
}

// Subnet is a DHCP subnet.
type Subnet struct {
	Network string
	Netmask string
	Range   Range
}

// Range is a range in which IP address should be given out.
type Range struct {
	Start string
	End   string
}

// Configure configures the DHCP server.
func (d *DHCPD) Configure() error {
	var configFileContent string
	for _, subnet := range d.Subnets {
		header := fmt.Sprintf("subnet %s netmask %s {", subnet.Network, subnet.Netmask)

		ranges := fmt.Sprintf("range %s %s;", subnet.Range.Start, subnet.Range.End)

		footer := "}"

		configFileContent += fmt.Sprintf("%s\n\t%s\n%s\n", header, ranges, footer)
	}
	if err := os.MkdirAll(d.StateDir, os.ModePerm); err != nil {
		return err
	}

	d.configFileDir = filepath.Join(d.StateDir, "dhcpd.conf")
	configFile, err := os.Create(d.configFileDir)
	if err != nil {
		return err
	}

	if _, err = configFile.WriteString(configFileContent); err != nil {
		return err
	}
	defer configFile.Close()

	d.leasesFileDir = filepath.Join(d.StateDir, "dhcpd.leases")
	leasesFile, err := os.Create(d.leasesFileDir)
	if err != nil {
		return err
	}
	defer leasesFile.Close()

	fmt.Println(d.configFileDir, d.leasesFileDir)

	return nil
}
