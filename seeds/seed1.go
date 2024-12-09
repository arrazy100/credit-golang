package main

import (
	"credit/config"
	"credit/dtos/request"
	"credit/models"
	"credit/models/base"
	"credit/models/enums"
	"credit/services"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Seed1(config *config.Config) error {
	authService := services.NewAuthService(config.DatabaseConnection)

	registerPayloads := []request.RegisterPayload{
		{
			Email:    "budi@user.com",
			Password: "budi123",
		},
		{
			Email:    "annisa@user.com",
			Password: "annisa123",
		},
	}

	var registeredUserID []uuid.UUID
	for _, payload := range registerPayloads {
		response, _, err := authService.RegisterUser(payload)
		if err != nil {
			return fmt.Errorf("cannot register user with email %s", payload.Email)
		}

		registeredUserID = append(registeredUserID, response.UserID)
	}

	debtors := []models.Debtor{
		{
			BaseUser: base.BaseUser{
				UserID: registeredUserID[0],
			},
			BaseAuditUser: base.BaseAuditUser{
				CreatedByID: registeredUserID[0],
				BaseAudit: base.BaseAudit{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
				},
			},
			NIK:                "3232323232323231",
			FullName:           "Budi",
			LegalName:          "Budi",
			PlaceOfBirth:       "Jakarta",
			DateOfBirth:        time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Salary:             "6000000",
			IdentityPictureUrl: "http://identity-budi.com",
			SelfiePictureUrl:   "http://selfie-budi.com",
		},
		{
			BaseUser: base.BaseUser{
				UserID: registeredUserID[1],
			},
			BaseAuditUser: base.BaseAuditUser{
				CreatedByID: registeredUserID[1],
				BaseAudit: base.BaseAudit{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
				},
			},
			NIK:                "3232323232323232",
			FullName:           "Annisa",
			LegalName:          "Annisa",
			PlaceOfBirth:       "Jakarta",
			DateOfBirth:        time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Salary:             "10000000",
			IdentityPictureUrl: "http://identity-annisa.com",
			SelfiePictureUrl:   "http://selfie-annisa.com",
		},
	}

	tenorLimitBudi := []models.DebtorTenorLimit{
		{
			BaseUser: base.BaseUser{
				UserID: registeredUserID[0],
			},
			BaseAuditUser: base.BaseAuditUser{
				CreatedByID: registeredUserID[0],
				BaseAudit: base.BaseAudit{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
				},
			},
			DebtorID:       debtors[0].ID,
			TenorLimitType: enums.Monthly,
			TenorDuration:  1,
			TotalLimit:     "100000",
			CurrentLimit:   "100000",
		},
		{
			BaseUser: base.BaseUser{
				UserID: registeredUserID[0],
			},
			BaseAuditUser: base.BaseAuditUser{
				CreatedByID: registeredUserID[0],
				BaseAudit: base.BaseAudit{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
				},
			},
			DebtorID:       debtors[0].ID,
			TenorLimitType: enums.Monthly,
			TenorDuration:  2,
			TotalLimit:     "200000",
			CurrentLimit:   "200000",
		},
		{
			BaseUser: base.BaseUser{
				UserID: registeredUserID[0],
			},
			BaseAuditUser: base.BaseAuditUser{
				CreatedByID: registeredUserID[0],
				BaseAudit: base.BaseAudit{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
				},
			},
			DebtorID:       debtors[0].ID,
			TenorLimitType: enums.Monthly,
			TenorDuration:  3,
			TotalLimit:     "500000",
			CurrentLimit:   "500000",
		},
		{
			BaseUser: base.BaseUser{
				UserID: registeredUserID[0],
			},
			BaseAuditUser: base.BaseAuditUser{
				CreatedByID: registeredUserID[0],
				BaseAudit: base.BaseAudit{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
				},
			},
			DebtorID:       debtors[0].ID,
			TenorLimitType: enums.Monthly,
			TenorDuration:  6,
			TotalLimit:     "700000",
			CurrentLimit:   "700000",
		},
	}

	debtors[0].TenorLimits = tenorLimitBudi

	tenorLimitAnnisa := []models.DebtorTenorLimit{
		{
			BaseUser: base.BaseUser{
				UserID: registeredUserID[1],
			},
			BaseAuditUser: base.BaseAuditUser{
				CreatedByID: registeredUserID[1],
				BaseAudit: base.BaseAudit{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
				},
			},
			DebtorID:       debtors[0].ID,
			TenorLimitType: enums.Monthly,
			TenorDuration:  1,
			TotalLimit:     "1000000",
			CurrentLimit:   "1000000",
		},
		{
			BaseUser: base.BaseUser{
				UserID: registeredUserID[1],
			},
			BaseAuditUser: base.BaseAuditUser{
				CreatedByID: registeredUserID[1],
				BaseAudit: base.BaseAudit{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
				},
			},
			DebtorID:       debtors[0].ID,
			TenorLimitType: enums.Monthly,
			TenorDuration:  2,
			TotalLimit:     "1200000",
			CurrentLimit:   "1200000",
		},
		{
			BaseUser: base.BaseUser{
				UserID: registeredUserID[1],
			},
			BaseAuditUser: base.BaseAuditUser{
				CreatedByID: registeredUserID[1],
				BaseAudit: base.BaseAudit{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
				},
			},
			DebtorID:       debtors[0].ID,
			TenorLimitType: enums.Monthly,
			TenorDuration:  3,
			TotalLimit:     "1500000",
			CurrentLimit:   "1500000",
		},
		{
			BaseUser: base.BaseUser{
				UserID: registeredUserID[1],
			},
			BaseAuditUser: base.BaseAuditUser{
				CreatedByID: registeredUserID[1],
				BaseAudit: base.BaseAudit{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
				},
			},
			DebtorID:       debtors[0].ID,
			TenorLimitType: enums.Monthly,
			TenorDuration:  6,
			TotalLimit:     "2000000",
			CurrentLimit:   "2000000",
		},
	}

	debtors[1].TenorLimits = tenorLimitAnnisa

	err := config.DatabaseConnection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&debtors).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
