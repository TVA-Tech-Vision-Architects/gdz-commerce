package repository

import (
	"github.com/B6137151/GDZ-Commerce/internal/domain/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindAdminByEmail(email string) (*model.Admin, error)
	FindUserByEmail(email string) (*model.User, error)
	CreateAdmin(admin *model.Admin) error
	CreateUser(user *model.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) FindAdminByEmail(email string) (*model.Admin, error) {
	var admin model.Admin
	err := r.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *authRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) CreateAdmin(admin *model.Admin) error {
	return r.db.Create(admin).Error
}

func (r *authRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}
