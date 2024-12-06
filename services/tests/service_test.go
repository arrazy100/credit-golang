package services_tests

import (
	"credit/models/base"
	"credit/models/enums"
	"credit/utils"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ServiceTestSuite struct {
	suite.Suite
	Db    *gorm.DB
	Admin base.User
}

func (suite *ServiceTestSuite) SetupTest() {
	var err error
	suite.Db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	suite.Db.AutoMigrate(&base.User{})

	SetupAdminUser(suite, &suite.Suite)
}

func (suite *ServiceTestSuite) TearDownTest() {
	sqlDB, err := suite.Db.DB()
	if err != nil {
		panic("failed to get SQL DB")
	}
	sqlDB.Close()
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func SetupAdminUser(suiteStruct *ServiceTestSuite, suite *suite.Suite) {
	hashed, err := utils.HashPassword("admin")
	suite.NoError(err)

	suiteStruct.Admin = base.User{
		ID:       uuid.New(),
		Email:    "admin@admin.com",
		Password: hashed,
		Role:     enums.Admin,
	}

	err = suiteStruct.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&suiteStruct.Admin).Error
		suite.NoError(err)

		return nil
	})
	suite.NoError(err)
}
