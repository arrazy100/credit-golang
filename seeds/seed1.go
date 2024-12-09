package main

import (
	"credit/config"
	"credit/dtos/request"
	"credit/services"
	"fmt"

	"github.com/google/uuid"
)

func Seed1(config *config.Config) error {
	debtorService := services.NewDebtorService(config.DatabaseConnection)
	authService := services.NewAuthService(config.DatabaseConnection)

	registerPayloads := []request.RegisterPayload{
		{
			Email:    "andi@user.com",
			Password: "andi123",
		},
		{
			Email:    "budi@user.com",
			Password: "budi123",
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

	registerDebtorPayloads := []request.RegisterDebtorPayload{
		{
			NIK:                "3232323232323231",
			FullName:           "Andi Pratama",
			LegalName:          "Andi Pratama",
			PlaceOfBirth:       "Jakarta",
			DateOfBirth:        "1990-03-03",
			Salary:             "6000000",
			IdentityPictureUrl: "http://identity-andi.com",
			SelfiePictureUrl:   "http://selfie-andi.com",
		},
		{
			NIK:                "3232323232323232",
			FullName:           "Budi Pangestu",
			LegalName:          "Budi Pangestu",
			PlaceOfBirth:       "Jakarta",
			DateOfBirth:        "1990-03-04",
			Salary:             "10000000",
			IdentityPictureUrl: "http://identity-budi.com",
			SelfiePictureUrl:   "http://selfie-budi.com",
		},
	}

	for i, payload := range registerDebtorPayloads {
		_, _, err := debtorService.Register(registeredUserID[i], payload)
		if err != nil {
			return fmt.Errorf("cannot register Debtor with userID %s", registeredUserID[i])
		}
	}

	return nil
}
