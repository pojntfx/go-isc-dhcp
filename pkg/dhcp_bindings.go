package pkg

//go:generate sh -c "rm -rf dhcp; git clone https://gitlab.isc.org/isc-projects/dhcp.git; cd dhcp; ./configure; make"

/*
#cgo LDFLAGS: -ldhcp -lomapi -ldhcpctl -lirs -ldns -lisc -lisccfg -Ldhcp/bind/bind-9.11.14/lib/isc -Ldhcp/bind/bind-9.11.14/lib/isccfg -Ldhcp/bind/bind-9.11.14/lib/dns -Ldhcp/bind/bind-9.11.14/lib/irs -Ldhcp/dhcpctl -L dhcp/omapip -Ldhcp/common
#cgo CFLAGS: -w -Idhcp/includes -Idhcp/includes/omapip -Idhcp/bind/include -Idhcp/bind/bind-9.11.14/lib/isc/include/isc

#include "dhcpd_bindings.h"
*/
import "C"
import (
	"errors"
	"fmt"
)

// DHCPServer is an ISC DHCP server.
type DHCPServer struct {
	ConfigurationFile string
	Device            string
}

// Start starts the DHCP server.
func (s *DHCPServer) Start() error {
	if rv := int(C.main_extracted(C.CString(s.ConfigurationFile), C.CString(s.Device))); rv != 0 {
		return errors.New("could not start dhcp server, exit code " + fmt.Sprintf("%v", rv))
	}

	return nil
}
