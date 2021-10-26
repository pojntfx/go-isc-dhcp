package workers

import (
	_ "embed"
	"os"
	"os/exec"
	"syscall"

	"github.com/pojntfx/go-isc-dhcp/pkg/utils"
)

var (
	//go:embed dhcp/client/dhclient
	EmbeddedDHClient []byte
)

// DHClient is a dhcp client.
type DHClient struct {
	utils.ProcessWorker
	Device    string
	BinaryDir string
}

// Start starts the the dhcp client.
func (d *DHClient) Start() error {
	d.ScheduledForDeletion = false

	command := exec.Command(d.BinaryDir, "-d", "-i", d.Device)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	d.Instance = command

	return command.Start()
}
