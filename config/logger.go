package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/spf13/viper"
	"os"
	"sync"
	"time"
)

var (
	once   = new(sync.Once)
	Logger zerolog.Logger
)

func InitLogger() {
	once.Do(func() {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

		if viper.GetString("server_mode") == "debug" {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}

		Logger = zerolog.New(diode.NewWriter(os.Stdout, 10000, 10*time.Millisecond, func(missed int) {
			fmt.Printf("Dropped %d messages", missed)
		}))
	})
}
