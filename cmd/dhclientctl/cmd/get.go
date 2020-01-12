package cmd

import (
	"context"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/gosuri/uitable"
	constants "github.com/pojntfx/go-isc-dhcp/cmd"
	goISCDHCP "github.com/pojntfx/go-isc-dhcp/pkg/proto/generated"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
)

var getCmd = &cobra.Command{
	Use:     "get [id]",
	Aliases: []string{"g"},
	Short:   "Get one or all dhcp client(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(viper.GetString(serverHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := goISCDHCP.NewDHClientManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if len(args) < 1 {
			response, err := client.List(ctx, &goISCDHCP.DHClientManagerListArgs{})
			if err != nil {
				return err
			}

			table := uitable.New()
			table.AddRow(
				"ID",
				"DEVICE")

			for _, DHClient := range response.GetDHClientsManaged() {
				table.AddRow(
					DHClient.GetId(),
					DHClient.GetDevice())
			}

			fmt.Println(table)

			return nil
		}

		response, err := client.Get(ctx, &goISCDHCP.DHClientManagedId{
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
	var (
		serverHostPortFlag string
	)

	getCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.DHClientDHostPortDefault, "Host:port of the go-isc-dhcp server to use.")

	if err := viper.BindPFlags(getCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(getCmd)
}
