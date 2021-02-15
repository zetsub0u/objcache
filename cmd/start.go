package cmd

import (
	"os"
	"os/signal"
	"strings"

	"github.com/zetsub0u/objcache/manager"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zetsub0u/objcache/http"
)

var flagHelper viperPFlagHelper

var (
	addressFlag string
	portFlag    int
)

func printSettings() {
	log.Infof("Bind Address: %s", addressFlag)
	log.Infof("Bind Port: %d", portFlag)
}

func start() {
	printSettings()

	// wire signal handlers
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// http api server
	log.Info("cmd: initializing http api")
	apiConfig := http.ServerConfig{Address: addressFlag, Port: portFlag}
	apiServer := http.NewServer(&apiConfig).WithMetrics().WithManager(manager.NewObjectStore())
	apiServer.Setup()

	log.Info("cmd: starting http api...")
	go apiServer.Start()

	<-quit
	log.Info("cmd: warm shutdown initiated...")

	// on double signal exit immediately
	go func() {
		<-quit
		log.Info("cmd: cold shutdown requested, bye.")
		os.Exit(1)
	}()

	apiServer.Stop()

	log.Info("cmd: exiting...")
	os.Exit(0)
}

func init() {
	// set default prefix for environment variable overrides, nested variables use _ instead of . for sublevels
	viper.SetEnvPrefix("objcache")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// Base FLags
	startCmd.Flags().StringVarP(&addressFlag, "bind-addr", "b", "localhost", "bind address for the web server")
	startCmd.Flags().IntVarP(&portFlag, "bind-port", "p", 8080, "bind port for the web server")

	// Add commands
	RootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the object cache rest api server",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}
