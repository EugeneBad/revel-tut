package models

import (
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Email 		string `sql:"unique"`
	Password 	string
}
