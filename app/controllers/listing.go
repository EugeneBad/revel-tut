package controllers

import (
	"github.com/revel/revel"
	"image"
	"github.com/eugenebad/massape/app/models"
	"bytes"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type Listing struct{
	*revel.Controller
}

func (l Listing) Create(pic []byte, ) revel.Result{

	switch l.Request.Method {
	case "GET":
		return l.Render()
	case "POST":
		var listing models.Listing
		l.Params.Bind(&listing, "listing")
		fmt.Println(l.Request.GetForm())
		f, _ := os.Create("picture")
		defer f.Close()
		f.Write(pic)
		image.DecodeConfig(bytes.NewReader(pic)) // For validation
		//return l.Render
	}
	return nil
}