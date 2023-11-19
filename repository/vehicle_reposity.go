package repository

import (
	"github.com/teerapon19/transaction-gorm/repository/model"

	"gorm.io/gorm"
)

type IVehicleRepository interface {
	IUseTransaction
	Create(name string) error
}

type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) IVehicleRepository {
	return &vehicleRepository{
		db: db,
	}
}

func (r *vehicleRepository) Create(name string) error {
	return r.db.Create(&model.Vehicle{
		Name: name,
	}).Error
}

func (r *vehicleRepository) SetDB(db *gorm.DB) {
	r.db = db
}
