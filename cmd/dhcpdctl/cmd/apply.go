package cmd

import (
	"context"
	"fmt"

	"github.com/ghodss/yaml"
	constants "github.com/pojntfx/go-isc-dhcp/cmd"
	api "github.com/pojntfx/go-isc-dhcp/pkg/api/proto/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
)

var applyCmd = &cobra.Command{
	Use:     "apply",
	Aliases: []string{"a"},
	Short:   "Apply a dhcp server",
	RunE: func(cmd *cobra.Command, args []string) error {
		var subnets []*api.Subnet

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

		client := api.NewDHCPDManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		response, err := client.Create(ctx, &api.DHCPD{
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
		deviceFlag  string
		subnetsFlag string
	)

	applyCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.DHCPDDHostPortDefault, constants.HostPortDocs)
	applyCmd.PersistentFlags().StringVarP(&configFileFlag, configFileKey, "f", configFileDefault, constants.ConfigurationFileDocs)
	applyCmd.PersistentFlags().StringVarP(&deviceFlag, deviceKey, "d", "edge0", "Device to bind to.")
	applyCmd.PersistentFlags().StringVarP(&subnetsFlag, subnetsKey, "n", `[
  {
    "netmask": "255.255.255.0",
    "network": "192.168.1.0",
    "nextServer": "192.168.1.1",
	"filename": "undionly.kpxe",
	"routers": "192.168.178.1",
	"domainNameServers": ["8.8.8.8"],
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
