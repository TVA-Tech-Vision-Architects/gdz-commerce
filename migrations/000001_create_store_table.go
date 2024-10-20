package migrations

import (
	"github.com/B6137151/GDZ-Commerce/internal/domain/model"
	"gorm.io/gorm"
)

func CreateStoreTable(db *gorm.DB) error {
	// AutoMigrate for Store, Admin, and User models
	return db.AutoMigrate(
		&model.Store{},
		&model.Admin{},
		&model.User{},
	)
}
