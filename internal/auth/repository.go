package auth

import (
	"github.com/tafo/rosa/internal"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

var Repo AccountRepository

func InitAccountRepository(db *gorm.DB) {
	Repo.db = db
}

func (repo AccountRepository) Insert(account *Account) error {
	result := repo.db.Create(account)
	if err := result.Error; err != nil {
		internal.Logger.Error().Err(err).Msg("Account could be created")
		return err
	}
	return nil
}