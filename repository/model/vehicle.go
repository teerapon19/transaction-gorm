package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID   string
	Name string
}

func (Vehicle) TableName() string {
	return "fighter.vehicle"
}

func (u *Vehicle) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewString()
	return nil
}
