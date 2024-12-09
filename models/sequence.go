package models

type Sequence struct {
	ID         string `json:"id" gorm:"primary_key;varchar(255)"`
	Prefix     string `json:"prefix" gorm:"varchar(255)"`
	LastNumber int64  `json:"last_number" gorm:"not null"`
}
