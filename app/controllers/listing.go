package controllers

import (
	"github.com/revel/revel"
	"image"
	"github.com/eugenebad/massape/app/models"
	"bytes"
	_ "image/jpeg"
	_ "image/png"
	//"os"
	"strings"
	//"github.com/eugenebad/massape/app"
	"fmt"
	"github.com/eugenebad/massape/app/utils"
	"os"
)

type Listing struct{
	*revel.Controller
}

func (l Listing) Create(listingImage []byte, ) revel.Result{

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
		//app.Db.Create(&models.Listing{Title: listing.Title, Category:listing.Category, Description:listing.Description, ImageUrl:imageUrl, AccountID:1})
		f, _ := os.Create(imageUrl)
		defer f.Close()
		f.Write(listingImage)
		fmt.Println(imageUrl)
		return l.RenderText("Done")
	}
	return nil
}