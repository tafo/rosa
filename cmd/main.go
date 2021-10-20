package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tafo/rosa/config"
	"net/http"
	"os"
)

func init() {
	config.InitConfiguration()
	config.InitLogger()
	InitContainer()
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
	config.Logger.Info().Msg(fmt.Sprintf("Starting rosa (v%s) on :%d", viper.GetString("version"), viper.GetInt("server_port")))

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", viper.GetInt("server_port")),
	}

	err := server.ListenAndServe()
	if err != nil {
		config.Logger.Fatal().Err(err).Msg("Could not start the service")
	}

	exitCode := gracefullyShutdown(server)

	return exitCode, nil // Fix this
}