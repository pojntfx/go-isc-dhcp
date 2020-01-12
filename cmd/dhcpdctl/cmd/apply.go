package cmd

import (
	"context"
	"fmt"
	"github.com/ghodss/yaml"
	constants "github.com/pojntfx/godhcpd/cmd"
	godhcpd "github.com/pojntfx/godhcpd/pkg/proto/generated"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
)

var applyCmd = &cobra.Command{
	Use:     "apply",
	Aliases: []string{"a"},
	Short:   "Apply an dhcp server",
	RunE: func(cmd *cobra.Command, args []string) error {
		var subnets []*godhcpd.Subnet

		if !(viper.GetString(configFileKey) == configFileDefault) {
			viper.SetConfigFile(viper.GetString(configFileKey))

			if err := viper.ReadInConfig(); err != nil {
				return err
			}

			if err := viper.UnmarshalKey(subnetsKey, &subnets); err != nil {
				return err
			}
		} else {
			if err := yaml.Unmarshal([]byte(viper.GetString(subnetsKey)), &subnets); err != nil {
				return err
			}
		}

		conn, err := grpc.Dial(viper.GetString(serverHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := godhcpd.NewDHCPDManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		response, err := client.Create(ctx, &godhcpd.DHCPD{
			Device:  viper.GetString(deviceKey),
			Subnets: subnets,
		})
		if err != nil {
			return err
		}

		fmt.Printf("dhcp server \"%s\" created\n", response.GetId())

		return nil
	},
}

func init() {
	var (
		serverHostPortFlag string
		configFileFlag     string
		deviceFlag         string
		subnetsFlag        string
	)

	applyCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.DHCPDDHostPortDefault, "Host:port of the godhcpd server to use.")
	applyCmd.PersistentFlags().StringVarP(&configFileFlag, configFileKey, "f", configFileDefault, "Configuration file to use.")
	applyCmd.PersistentFlags().StringVarP(&deviceFlag, deviceKey, "d", "edge0", "Device to bind to.")
	applyCmd.PersistentFlags().StringVarP(&subnetsFlag, subnetsKey, "n", `[
  {
    "netmask": "255.255.255.0",
    "network": "192.168.1.0",
    "range": {
      "start": "192.168.1.10",
      "end": "192.168.1.100"
    }
  }
]`, "Subnet declaration.")

	if err := viper.BindPFlags(applyCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(applyCmd)
}
