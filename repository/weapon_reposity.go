package repository

import (
	"github.com/teerapon19/transaction-gorm/repository/model"

	"gorm.io/gorm"
)

type IWeaponRepository interface {
	IUseTransaction
	Create(name string) error
}

type weaponRepository struct {
	db *gorm.DB
}

func NewWeaponReposity(db *gorm.DB) IWeaponRepository {
	return &weaponRepository{
		db: db,
	}
}

func (r *weaponRepository) Create(name string) error {
	return r.db.Create(&model.Weapon{
		Name: name,
	}).Error
}

func (r *weaponRepository) SetDB(db *gorm.DB) {
	r.db = db
}
