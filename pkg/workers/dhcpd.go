package workers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

// DHCPD is a DHCP server.
type DHCPD struct {
	Subnets                []Subnet
	BinaryDir              string
	ID                     string
	StateDir               string
	Device                 string
	configFileDir          string
	leasesFileDir          string
	instance               *exec.Cmd
	isScheduledForDeletion bool
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

	return nil
}

// Start starts the the DHCP server.
func (d *DHCPD) Start() error {
	d.isScheduledForDeletion = false

	command := exec.Command(d.BinaryDir, "-f", "-cf", d.configFileDir, "-lf", d.leasesFileDir, d.Device)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	d.instance = command

	return command.Start()
}

// Wait waits for the DHCP server to stop.
func (d *DHCPD) Wait() error {
	_, err := d.instance.Process.Wait()

	return err
}

// DisableAutoRestart disables the auto restart if the DHCP server exits.
func (d *DHCPD) DisableAutoRestart() error {
	d.isScheduledForDeletion = true

	return nil
}

// Stop stops the DHCP server.
func (d *DHCPD) Stop() error {
	if err := d.DisableAutoRestart(); err != nil {
		return err
	}

	processGroupID, err := syscall.Getpgid(d.instance.Process.Pid)
	if err != nil {
		return err
	}

	if err := syscall.Kill(processGroupID, syscall.SIGKILL); err != nil {
		return err
	}

	return nil
}

// IsScheduledForDeletion returns true if the DHCP server is scheduled for deletion.
func (d *DHCPD) IsScheduledForDeletion() bool {
	return d.isScheduledForDeletion
}

// IsRunning returns true if the DHCP server is still running.
func (d *DHCPD) IsRunning() bool {
	return d.instance.Process != nil
}

// Cleanup deletes the state of the DHCP server.
func (d *DHCPD) Cleanup() error {
	return os.RemoveAll(d.StateDir)
}
