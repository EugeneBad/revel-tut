package models

import (
	"github.com/jinzhu/gorm"
)

type Listing struct{
	gorm.Model
	Category string
	Title string
	Description string `sql:"size:600"`
}
