package repository

import (
	"github.com/B6137151/GDZ-Commerce/internal/domain/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StoreRepository interface {
	Create(store *model.Store) error
	GetByID(id uuid.UUID) (*model.Store, error)
	Update(store *model.Store) error
	Delete(id uuid.UUID) error
}

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &storeRepository{db: db}
}

func (r *storeRepository) Create(store *model.Store) error {
	return r.db.Create(store).Error
}

func (r *storeRepository) GetByID(id uuid.UUID) (*model.Store, error) {
	var store model.Store
	err := r.db.First(&store, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *storeRepository) Update(store *model.Store) error {
	return r.db.Save(store).Error
}

func (r *storeRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.Store{}, "id = ?", id).Error
}
