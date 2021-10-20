package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tafo/rosa/config"
	"github.com/tafo/rosa/internal"
	"os"
)

func init() {
	config.InitConfiguration()
	internal.InitLogger()
}

func main() {
	exit, err := run()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		exit = 1
	}

	// Exit with success status.
	os.Exit(exit)
}
func run() (int, error) {
	internal.Logger.Info().Msg(fmt.Sprintf("Starting rosa (v%s) on :%d", viper.GetString("version"), viper.GetInt("server_port")))

	server := CreateWebServer()

	internal.Logger.Info().Msg(fmt.Sprintf("Hede %s", server.Addr))
	err := server.ListenAndServe()
	if err != nil {
		internal.Logger.Fatal().Err(err).Msg("Could not start the service")
	}

	exitCode := gracefullyShutdown(server)

	return exitCode, nil // Fix this
}