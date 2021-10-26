package workers

//go:generate bash -c "rm -rf dhcp && git clone --depth 1 https://gitlab.isc.org/isc-projects/dhcp.git && cd dhcp && ./configure && make"

import (
	_ "embed"
	"fmt"
	"os"

	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/pojntfx/go-isc-dhcp/pkg/utils"
)

var (
	//go:embed dhcp/server/dhcpd
	EmbeddedDHCPD []byte
)

// DHCPD is a dhcp server.
type DHCPD struct {
	utils.ProcessWorker
	Subnets       []Subnet
	BinaryDir     string
	ID            string
	StateDir      string
	Device        string
	configFileDir string
	leasesFileDir string
}

// Subnet is a dhcp subnet.
type Subnet struct {
	Network           string
	Netmask           string
	NextServer        string
	Filename          string
	Routers           string
	DomainNameServers []string
	Range             Range
}

// Range is a range in which IP address should be given out.
type Range struct {
	Start string
	End   string
}

// Configure configures the dhcp server.
func (d *DHCPD) Configure() error {
	configFileContent := ""
	for _, subnet := range d.Subnets {
		configFileContent += fmt.Sprintf("subnet %s netmask %s {\n", subnet.Network, subnet.Netmask)

		configFileContent += fmt.Sprintf("\trange %s %s;\n", subnet.Range.Start, subnet.Range.End)

		if subnet.NextServer != "" {
			configFileContent += fmt.Sprintf("\tnext-server %s;\n", subnet.NextServer)
		}

		if subnet.Filename != "" {
			configFileContent += fmt.Sprintf("\tfilename \"%s\";\n", subnet.Filename)
		}

		if subnet.Routers != "" {
			configFileContent += fmt.Sprintf("\toption routers %s;\n", subnet.Routers)
		}

		if len(subnet.DomainNameServers) > 0 {
			domainNameServerLine := strings.Join(subnet.DomainNameServers, ",")

			configFileContent += fmt.Sprintf("\toption domain-name-servers %s;\n", domainNameServerLine)
		}

		configFileContent += "}\n"
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

	return nil
}

// Start starts the the dhcp server.
func (d *DHCPD) Start() error {
	d.ScheduledForDeletion = false

	command := exec.Command(d.BinaryDir, "-f", "-cf", d.configFileDir, "-lf", d.leasesFileDir, d.Device)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	d.Instance = command

	return command.Start()
}

// Cleanup deletes the state of the dhcp server.
func (d *DHCPD) Cleanup() error {
	return os.RemoveAll(d.StateDir)
}
