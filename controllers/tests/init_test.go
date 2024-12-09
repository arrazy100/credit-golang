package controllers_tests

import (
	"credit/dtos/request"
	"credit/dtos/response"
	"credit/models"
	"credit/models/base"
	"credit/models/enums"
	"credit/services"
	"credit/utils"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ControllerTestSuite struct {
	suite.Suite
	Db               *gorm.DB
	Service          *services.Service
	Admin            base.User
	DebtorUser       base.User
	TestUser         base.User
	AdminToken       string
	DebtorToken      string
	TestUserToken    string
	RegisteredDebtor response.RegisterDebtorResponse
}

func (suite *ControllerTestSuite) SetupTest() {
	var err error
	suite.Db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	suite.Db.AutoMigrate(
		&base.User{},
		&models.SeedVersion{},
		&models.Sequence{},
		&models.Debtor{},
		&models.DebtorTenorLimit{},
		&models.DebtorTransaction{},
		&models.DebtorInstallment{},
		&models.DebtorInstallmentLine{},
	)

	SetupService(suite, &suite.Suite)
	SetupUser(suite, &suite.Suite)
	SetupToken(suite, &suite.Suite)
	SetupDebtor(suite, &suite.Suite)
}

func (suite *ControllerTestSuite) TearDownTest() {
	sqlDB, err := suite.Db.DB()
	if err != nil {
		panic("failed to get SQL DB")
	}
	sqlDB.Close()
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}

func SetupUser(suiteStruct *ControllerTestSuite, suite *suite.Suite) {
	hashed, err := utils.HashPassword("password123")
	suite.NoError(err)

	suiteStruct.Admin = base.User{
		ID:       uuid.New(),
		Email:    "admin@admin.com",
		Password: hashed,
		Role:     enums.Admin,
	}

	suiteStruct.DebtorUser = base.User{
		ID:       uuid.New(),
		Email:    "debtor@debtor.com",
		Password: hashed,
		Role:     enums.Debtor,
	}

	suiteStruct.TestUser = base.User{
		ID:       uuid.New(),
		Email:    "test@debtor.com",
		Password: hashed,
		Role:     enums.Debtor,
	}

	err = suiteStruct.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&suiteStruct.Admin).Error
		suite.NoError(err)

		err = tx.Create(&suiteStruct.DebtorUser).Error
		suite.NoError(err)

		err = tx.Create(&suiteStruct.TestUser).Error
		suite.NoError(err)

		return nil
	})
	suite.NoError(err)
}

func SetupService(suiteStruct *ControllerTestSuite, suite *suite.Suite) {
	suiteStruct.Service = services.NewService(suiteStruct.Db)
}

func SetupToken(suiteStruct *ControllerTestSuite, suite *suite.Suite) {
	adminPayload := request.LoginPayload{
		Email:    suiteStruct.Admin.Email,
		Password: "password123",
	}
	adminResponse, _, _ := suiteStruct.Service.AuthService.Login(adminPayload)

	suiteStruct.AdminToken = "Bearer " + adminResponse.Token

	debtorPayload := request.LoginPayload{
		Email:    suiteStruct.DebtorUser.Email,
		Password: "password123",
	}
	debtorResponse, _, _ := suiteStruct.Service.AuthService.Login(debtorPayload)

	suiteStruct.DebtorToken = "Bearer " + debtorResponse.Token

	testUserPayload := request.LoginPayload{
		Email:    suiteStruct.TestUser.Email,
		Password: "password123",
	}
	testUserResponse, _, _ := suiteStruct.Service.AuthService.Login(testUserPayload)

	suiteStruct.TestUserToken = "Bearer " + testUserResponse.Token
}

func SetupDebtor(suiteStruct *ControllerTestSuite, suite *suite.Suite) {
	debtorPayload := request.RegisterDebtorPayload{
		NIK:                "3232323232323231",
		FullName:           "Test",
		LegalName:          "Test",
		PlaceOfBirth:       "Jakarta",
		DateOfBirth:        "1990-03-03",
		Salary:             "6000000",
		IdentityPictureUrl: "http://identity-test.com",
		SelfiePictureUrl:   "http://selfie-test.com",
	}

	registeredDebtor, _, _ := suiteStruct.Service.DebtorService.Register(suiteStruct.TestUser.ID, debtorPayload)

	suiteStruct.RegisteredDebtor = *registeredDebtor
}
