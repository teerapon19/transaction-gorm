package main

import (
	"github.com/teerapon19/transaction-gorm/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=12345678 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	vehicleRepo := repository.NewVehicleRepository(db)
	weaponRepo := repository.NewWeaponReposity(db)
	transactionRepo := repository.NewTransaction(db, struct {
		UserRepo    repository.IUserRepository
		VehicleRepo repository.IVehicleRepository
		WeaponRepo  repository.IWeaponRepository
	}{
		UserRepo:    userRepo,
		VehicleRepo: vehicleRepo,
		WeaponRepo:  weaponRepo,
	})

	err = transactionRepo.DoTransaction(func(tx *gorm.DB, repo struct {
		UserRepo    repository.IUserRepository
		VehicleRepo repository.IVehicleRepository
		WeaponRepo  repository.IWeaponRepository
	}) error {
		if err := repo.UserRepo.Create("user"); err != nil {
			return err
		}
		if err := repo.WeaponRepo.Create("weapon"); err != nil {
			return err
		}
		if err := repo.VehicleRepo.Create("vehicle"); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

}
