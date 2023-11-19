package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID   string
	Name string
}

func (*User) TableName() string {
	return "fighter.user"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewString()
	return nil
}
