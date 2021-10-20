package config

import (
	"database/sql"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	db, err := sql.Open("postgres", viper.GetString("DATABASE_URL"))
	if err != nil {
		Logger.Fatal().Err(err).Msg("Db connection failed")
		return nil
	}

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})

	conn, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		Logger.Fatal().Err(err).Msg("Db connection failed")
		return nil
	}

	if err = db.Ping(); err != nil {
		Logger.Fatal().Err(err).Msg("Ping failed")
		return nil
	}

	return conn
}
