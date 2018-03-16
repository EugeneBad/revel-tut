package models

import (
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Email 		string `sql:"unique"`
	Password 	string
	Listings 	[]Listing
}

func (a *Account) HasID() bool{
	if a.ID != 0 {return true}
	return false
}