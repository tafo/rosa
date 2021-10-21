package auth

import (
	"github.com/tafo/rosa/internal"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

var Repo Repository

func InitRepository(db *gorm.DB) {
	Repo.db = db
}

func (repo Repository) Insert(account *Account) error {
	result := repo.db.Create(account)
	if err := result.Error; err != nil {
		internal.Logger.Error().Err(err).Msg("Account could be created")
		return err
	}
	return nil
}