package svc

//go:generate sh -c "rm -rf dhcp statik; git clone https://gitlab.isc.org/isc-projects/dhcp.git; cd dhcp; ./configure; make; mkdir -p dist; cp server/dhcpd dist; cp client/dhclient dist; statik -src dhcp/dist"

import (
	_ "github.com/pojntfx/godhcpd/pkg/svc/statik"
	"github.com/rakyll/statik/fs"
	"os"
)

type DHCPDManager struct {
	BinaryDir string
}

func (m *DHCPDManager) Extract() error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}

	data, err := fs.ReadFile(statikFS, "/dhcpd")
	if err != nil {
		return err
	}

	binaryFile, err := os.Create(m.BinaryDir)
	if err != nil {
		return err
	}

	if _, err = binaryFile.Write(data); err != nil {
		return err
	}
	defer binaryFile.Close()

	if err := os.Chmod(m.BinaryDir, os.ModePerm); err != nil {
		return err
	}

	return nil
}
