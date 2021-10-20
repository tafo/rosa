package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tafo/rosa/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func gracefullyShutdown(srv *http.Server) int {
	sig := make(chan os.Signal, 1)
	exit := 0
	defer close(sig)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Block and receive signal
	in := <-sig

	errCh := make(chan error)
	defer close(errCh)

	config.Logger.Info().Msg(fmt.Sprintf("Received signal: %s, shutting down", in.String()))

	go func(s *http.Server, errCh chan error) {
		timeoutDuration := time.Duration(viper.GetInt("graceful_shutdown_timeout_in_seconds")) * time.Second
		timeout, cancel := context.WithTimeout(context.Background(), timeoutDuration)
		defer cancel()

		// server shutdown
		if err := s.Shutdown(timeout); err != nil {
			errCh <- err
		}

		// other shutdowns

		errCh <- nil
	}(srv, errCh)

	// Receive message from shutdown
	e := <-errCh

	if e != nil {
		// EX_SOFTWARE. See "sysexits.h"
		exit = 70

		config.Logger.Error().Err(e).Msg("Failed to shutdown server gracefully")
	} else {
		config.Logger.Info().Msg("Successfully closed server")
	}

	// Terminate server
	config.Logger.Info().Msg("Shutdown complete")

	return exit
}
