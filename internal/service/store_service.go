package service

import (
	"github.com/B6137151/GDZ-Commerce/internal/domain/model"
	"github.com/B6137151/GDZ-Commerce/internal/domain/repository"
	"github.com/google/uuid"
)

type StoreService interface {
	CreateStore(storeName, location string) (*model.Store, error)
	GetStore(id uuid.UUID) (*model.Store, error)
	UpdateStore(id uuid.UUID, storeName, location string) (*model.Store, error)
	DeleteStore(id uuid.UUID) error
}

type storeService struct {
	repo repository.StoreRepository
}

func NewStoreService(repo repository.StoreRepository) StoreService {
	return &storeService{repo: repo}
}

func (s *storeService) CreateStore(storeName, location string) (*model.Store, error) {
	store := &model.Store{
		StoreName: storeName,
		Location:  location,
	}
	err := s.repo.Create(store)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s *storeService) GetStore(id uuid.UUID) (*model.Store, error) {
	return s.repo.GetByID(id)
}

func (s *storeService) UpdateStore(id uuid.UUID, storeName, location string) (*model.Store, error) {
	store, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	store.StoreName = storeName
	store.Location = location
	err = s.repo.Update(store)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s *storeService) DeleteStore(id uuid.UUID) error {
	return s.repo.Delete(id)
}
