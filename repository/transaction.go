package repository

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type IUseTransaction interface {
	SetDB(db *gorm.DB)
}

type ITransaction[T any] interface {
	DoTransaction(fc func(tx *gorm.DB, repo T) error) error
}

type transaction[T any] struct {
	db        *gorm.DB
	repoGroup T
}

func NewTransaction[N any](db *gorm.DB, repoGroup N) ITransaction[N] {
	return &transaction[N]{
		db:        db,
		repoGroup: repoGroup,
	}
}

func (t *transaction[T]) DoTransaction(fc func(tx *gorm.DB, repo T) error) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		tm := reflect.TypeOf(t.repoGroup)
		vm := reflect.ValueOf(t.repoGroup)

		newRepoGroup := reflect.New(tm).Elem()

		for i := 0; i < tm.NumField(); i++ {
			mainRepoType := vm.Field(i).Elem().Elem().Type()
			fieldRepo := reflect.New(mainRepoType)

			if _, ok := fieldRepo.Type().MethodByName("SetDB"); !ok {
				return fmt.Errorf("declare method SetDB is required")
			}

			fieldRepo.MethodByName("SetDB").Call([]reflect.Value{reflect.ValueOf(tx)})
			newRepoGroup.Field(i).Set(fieldRepo)
		}
		return fc(tx, newRepoGroup.Interface().(T))
	})
}
