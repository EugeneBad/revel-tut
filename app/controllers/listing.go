package controllers

import (
	"github.com/revel/revel"
	"image"
	"github.com/eugenebad/massape/app/models"
	"bytes"
	_ "image/jpeg"
	_ "image/png"
	"strings"
	"github.com/eugenebad/massape/app"
	"github.com/eugenebad/massape/app/utils"
	"os"
)

type Listing struct{
	*revel.Controller
}

func (l Listing) Create(listingImage []byte, ) revel.Result{
	if l.Session["email"] == ""{
		return l.Redirect(Accounts.Login)
	}
	switch l.Request.Method {
	case "GET":
		return l.Render()
	case "POST":
		var listing models.Listing
		l.Params.Bind(&listing, "listing")

		l.Validation.Required(listing.Title)
		l.Validation.Required(listing.Category)
		l.Validation.Required(listing.Description)

		if l.Validation.HasErrors(){
			return l.RenderText("All fields required")
		}

		_, format, err := image.DecodeConfig(bytes.NewReader(listingImage))

		if err != nil {
			return l.RenderText("Invalid image")
		}

		imageUrl := strings.Join([]string{"uploaded/", utils.Randomiser(), ".", format}, "")

		var current_user models.Account
		app.Db.Where("id = ?", 1).First(&current_user)

		new_listing := models.Listing{Title: listing.Title,
			Category:listing.Category,
			Description:listing.Description,
			ImageUrl: imageUrl}

		current_user.Listings =  append(current_user.Listings, new_listing)
		app.Db.Save(&current_user)

		f, _ := os.Create(imageUrl)
		defer f.Close()
		f.Write(listingImage)

		return l.RenderText("Done")
	}
	return nil
}

func (l *Listing) Listings() revel.Result{
	return l.Render()
}
