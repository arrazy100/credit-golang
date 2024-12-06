package main

import (
	"credit/config"
	"credit/models"
	"credit/models/base"
	"credit/models/enums"
	"credit/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedInitial(config *config.Config) error {
	hashed, err := utils.HashPassword("admin")
	if err != nil {
		return err
	}

	admin := base.User{
		ID:       uuid.New(),
		Email:    "admin@admin.com",
		Password: hashed,
		Role:     enums.Admin,
	}

	err = config.DatabaseConnection.Transaction(func(tx *gorm.DB) error {
		var seedVersion models.SeedVersion
		if err := tx.First(&seedVersion).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if seedVersion.Version < 1 {
			if err := tx.Create(&admin).Error; err != nil {
				return err
			}

			if err := tx.Create(&models.SeedVersion{Version: 1}).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
