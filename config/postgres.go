package config

import (
	"github.com/spf13/viper"
	"github.com/tafo/rosa/internal"
	"github.com/tafo/rosa/internal/auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	conn, err := gorm.Open(postgres.Open(viper.GetString("DB_DSN")), &gorm.Config{})
	_ = conn.AutoMigrate(&models.Account{})
	if err != nil {
		internal.Logger.Fatal().Err(err).Msg("Db connection failed")
		return nil
	}

	return conn
}
