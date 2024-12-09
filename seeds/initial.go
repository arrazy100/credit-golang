package main

import (
	"credit/config"
	"credit/models/base"
	"credit/models/enums"
	"credit/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedInitial(config *config.Config) error {
	hashed, err := utils.HashPassword("admin123")
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
		if err := tx.Create(&admin).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
