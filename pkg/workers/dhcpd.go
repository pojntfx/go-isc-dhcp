package workers

import (
	"fmt"
	"github.com/pojntfx/go-isc-dhcp/pkg/utils"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
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
	Network string
	Netmask string
	Range   Range
}

// Range is a range in which IP address should be given out.
type Range struct {
	Start string
	End   string
}

// Configure configures the dhcp server.
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
