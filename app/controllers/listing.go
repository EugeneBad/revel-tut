package controllers

import (
	"github.com/revel/revel"
	"image"
	"github.com/eugenebad/massape/app/models"
	"bytes"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
	"bytes"
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

		f, _ := os.Create(strings.Join([]string{"upload/", string(listing.ID)}, ""))

		defer f.Close()
		f.Write(listingImage)
		image.DecodeConfig(bytes.NewReader(listingImage))
		return l.Render
	}
	return nil
}