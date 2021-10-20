package auth

import (
	"github.com/tafo/rosa/internal"
	"github.com/tafo/rosa/internal/auth/models"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return AccountRepository{db: db}
}

func (repo AccountRepository) Insert(account *models.Account) error {
	result := repo.db.Create(account)
	if err := result.Error; err != nil {
		internal.Logger.Error().Err(err).Msg("Account could be created")
		return err
	}
	return nil
}