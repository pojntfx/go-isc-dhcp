package cmd

import (
	"context"
	"fmt"
	"sync"

	constants "github.com/pojntfx/go-isc-dhcp/cmd"
	api "github.com/pojntfx/go-isc-dhcp/pkg/api/proto/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
)

var deleteCmd = &cobra.Command{
	Use:     "delete <id> [id...]",
	Aliases: []string{"d"},
	Short:   "Delete one or more dhcp client(s)",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(viper.GetString(serverHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := api.NewDHClientManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var wg sync.WaitGroup

		for _, id := range args {
			wg.Add(1)

			go func(id string, wg *sync.WaitGroup) {
				response, err := client.Delete(ctx, &api.DHClientManagedId{
					Id: id,
				})
				if err != nil {
					log.Error(err.Error())

					wg.Done()

					return
				}

				fmt.Printf("dhcp client \"%s\" deleted\n", response.GetId())

				wg.Done()
			}(id, &wg)
		}

		wg.Wait()

		return nil
	},
}

func init() {
	deleteCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.DHClientDHostPortDefault, constants.HostPortDocs)

	if err := viper.BindPFlags(deleteCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(deleteCmd)
}
