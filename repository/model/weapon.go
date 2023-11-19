package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Weapon struct {
	ID   string
	Name string
}

func (Weapon) TableName() string {
	return "fighter.weapon"
}

func (u *Weapon) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewString()
	return nil
}
