package workers

import "fmt"

// DHCPD is a DHCP server.
type DHCPD struct {
	Subnets []Subnet
	ID      string
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
	var configFile string
	for _, subnet := range d.Subnets {
		header := fmt.Sprintf("subnet %s netmask %s {", subnet.Network, subnet.Netmask)

		ranges := fmt.Sprintf("range %s %s;", subnet.Range.Start, subnet.Range.End)

		footer := "}"

		configFile += fmt.Sprintf("%s\n\t%s\n%s\n", header, ranges, footer)
	}

	fmt.Println(configFile)

	return nil
}
