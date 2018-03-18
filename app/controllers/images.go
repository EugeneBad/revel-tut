package controllers

import (
	"github.com/revel/revel"
	"github.com/eugenebad/massape/app/models"
	"github.com/eugenebad/massape/app"
	"fmt"
	"os"
)

type Images struct {
	*revel.Controller
}

func (i Images) ImageServe() revel.Result {

	if i.Session["email"] == "" {
		return i.Forbidden("Login first")
	}
	var currentUser models.Account
	var currentListing models.Listing

	app.Db.Where("email = ?", i.Session["email"]).First(&currentUser)
	app.Db.Where("id = ? and account_id = ?", 7, currentUser.ID).First(&currentListing)
	f,_ := os.Open(currentListing.ImageUrl)
	fmt.Println(currentListing.ImageUrl)
	return i.RenderFile(f, revel.Inline)

}