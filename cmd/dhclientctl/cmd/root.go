package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "dhclientctl",
	Short: "dhclientctl manages dhclientd, the ISC DHCP client management daemon",
	Long: `dhclientctl manages dhclientd, the ISC DHCP client management daemon.

Find more information at:
https://pojntfx.github.io/go-isc-dhcp/`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("dhcpd")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	},
}

// Execute starts the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Could not start root command", rz.Err(err))
	}
}
