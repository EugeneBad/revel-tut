package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	type Phone struct {
		Kikka string
	}
	phone := Phone{Kikka: "XXXXXX"}

	return c.Render(phone)
}
