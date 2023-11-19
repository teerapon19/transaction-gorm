package repository

import (
	"github.com/teerapon19/transaction-gorm/repository/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	IUseTransaction
	Create(name string) error
	GetDB() *gorm.DB
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *userRepository) Create(name string) error {
	return r.db.Create(&model.User{
		Name: name,
	}).Error
}

func (r *userRepository) SetDB(db *gorm.DB) {
	r.db = db
}
