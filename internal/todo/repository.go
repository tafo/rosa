package todo

import (
	"github.com/tafo/rosa/internal"
	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

var Repo ItemRepository

func InitItemRepository(db *gorm.DB) {
	Repo.db = db
}

func (repo ItemRepository) GetAll(items *[]Item) error {
	result := repo.db.Find(items)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo ItemRepository) Insert(item *Item) error {
	result := Repo.db.Create(item)
	if result.Error != nil {
		internal.Logger.Error().Err(result.Error).Msg("Item could not be inserted")
		return result.Error
	}
	return nil
}

func (repo ItemRepository) Complete(item *Item) error {
	result := Repo.db.Model(&item).Update("is_completed", true)
	if result.Error != nil {
		internal.Logger.Error().Err(result.Error).Msg("Item could not be retrieved")
		return result.Error
	}
	return nil
}