package cmd

import (
	"context"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/gosuri/uitable"
	constants "github.com/pojntfx/go-isc-dhcp/cmd"
	api "github.com/pojntfx/go-isc-dhcp/pkg/api/proto/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
)

var getCmd = &cobra.Command{
	Use:     "get [id]",
	Aliases: []string{"g"},
	Short:   "Get one or all dhcp server(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(viper.GetString(serverHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := api.NewDHCPDManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if len(args) < 1 {
			response, err := client.List(ctx, &api.DHCPDManagerListArgs{})
			if err != nil {
				return err
			}

			table := uitable.New()
			table.AddRow(
				"ID",
				"DEVICE")

			for _, DHCPD := range response.GetDHCPDsManaged() {
				table.AddRow(
					DHCPD.GetId(),
					DHCPD.GetDevice())
			}

			fmt.Println(table)

			return nil
		}

		response, err := client.Get(ctx, &api.DHCPDManagedId{
			Id: args[0],
		})
		if err != nil {
			return err
		}

		output, err := yaml.Marshal(&response)
		if err != nil {
			return err
		}

		fmt.Println(string(output))

		return nil
	},
}

func init() {
	getCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.DHCPDDHostPortDefault, constants.HostPortDocs)

	if err := viper.BindPFlags(getCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(getCmd)
}
