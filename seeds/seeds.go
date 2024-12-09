package main

import (
	"credit/config"
	"credit/models"
	"log"

	"gorm.io/gorm"
)

func main() {
	configs, err := config.Load("config.dev.yaml")
	if err != nil {
		panic(err)
	}

	var seedVersion models.SeedVersion
	if err := configs.DatabaseConnection.First(&seedVersion).Error; err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}

	if seedVersion.Version < 1 {
		err = configs.DatabaseConnection.Transaction(func(tx *gorm.DB) error {
			err := SeedInitial(configs)

			seedVersion.Version++
			if err := tx.Save(&seedVersion).Error; err != nil {
				return err
			}

			return err
		})
		if err != nil {
			panic(err)
		}
	}

	if seedVersion.Version < 2 {
		err = configs.DatabaseConnection.Transaction(func(tx *gorm.DB) error {
			err := Seed1(configs)

			seedVersion.Version++
			if err := tx.Save(&seedVersion).Error; err != nil {
				return err
			}

			return err
		})
		if err != nil {
			panic(err)
		}
	}

	log.Println("Finished seeding")
}
