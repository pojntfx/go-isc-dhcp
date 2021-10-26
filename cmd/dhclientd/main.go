package main

import (
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	constants "github.com/pojntfx/go-isc-dhcp/cmd"
	api "github.com/pojntfx/go-isc-dhcp/pkg/api/proto/v1"
	"github.com/pojntfx/go-isc-dhcp/pkg/services"
	"github.com/pojntfx/go-isc-dhcp/pkg/workers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	keyPrefix         = "dhclientd."
	configFileDefault = ""
	configFileKey     = keyPrefix + "configFile"
	listenHostPortKey = keyPrefix + "listenHostPort"
)

var rootCmd = &cobra.Command{
	Use:   "dhclientd",
	Short: "dhclientd is the ISC DHCP client management daemon",
	Long: `dhclientd is the ISC DHCP client management daemon.

Find more information at:
https://github.com/pojntfx/go-isc-dhcp`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("dhclientd")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(viper.GetString(configFileKey) == configFileDefault) {
			viper.SetConfigFile(viper.GetString(configFileKey))

			if err := viper.ReadInConfig(); err != nil {
				return err
			}
		}
		binaryDir := filepath.Join(os.TempDir(), "dhclient")

		listener, err := net.Listen("tcp", viper.GetString(listenHostPortKey))
		if err != nil {
			return err
		}

		server := grpc.NewServer()
		reflection.Register(server)

		DHClientService := services.DHClientManager{
			BinaryDir:        binaryDir,
			DHClientsManaged: make(map[string]*workers.DHClient),
		}

		api.RegisterDHClientManagerServer(server, &DHClientService)

		interrupt := make(chan os.Signal, 2)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-interrupt

			// Allow manually killing the process
			go func() {
				<-interrupt

				os.Exit(1)
			}()

			log.Info("Gracefully stopping server (this might take a few seconds)")

			msg := "Could not stop dhcp server"

			for _, DHClient := range DHClientService.DHClientsManaged {
				if err := DHClient.DisableAutoRestart(); err != nil { // Manually disable auto restart; disables crash recovery even if process is not running
					log.Fatal(msg, rz.Err(err))
				}

				if DHClient.IsRunning() {
					if err := DHClient.Stop(); err != nil { // Stop is sync, so no need to `.Wait()`
						log.Fatal(msg, rz.Err(err))
					}
				}
			}

			if err := DHClientService.Cleanup(); err != nil {
				log.Fatal(msg, rz.Err(err))
			}

			server.GracefulStop()
		}()

		if err := DHClientService.Extract(); err != nil {
			return err
		}

		log.Info("Starting server")

		if err := server.Serve(listener); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	var (
		configFileFlag string
		hostPortFlag   string
	)

	rootCmd.PersistentFlags().StringVarP(&configFileFlag, configFileKey, "f", configFileDefault, constants.ConfigurationFileDocs)
	rootCmd.PersistentFlags().StringVarP(&hostPortFlag, listenHostPortKey, "l", constants.DHClientDHostPortDefault, "TCP listen host:port.")

	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(constants.CouldNotStartRootCommandErrorMessage, rz.Err(err))
	}
}
