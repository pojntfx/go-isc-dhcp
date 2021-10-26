package cmd

import (
	"strings"

	constants "github.com/pojntfx/go-isc-dhcp/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
)

var rootCmd = &cobra.Command{
	Use:   "dhcpdctl",
	Short: "dhcpdctl manages dhcpdd, the ISC DHCP server management daemon",
	Long: `dhcpdctl manages dhcpdd, the ISC DHCP server management daemon.

Find more information at:
https://github.com/pojntfx/go-isc-dhcp`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("dhcpd")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	},
}

// Execute starts the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(constants.CouldNotStartRootCommandErrorMessage, rz.Err(err))
	}
}
