package workers

import (
	"github.com/pojntfx/godhcpd/pkg/utils"
	"os"
	"os/exec"
	"syscall"
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
