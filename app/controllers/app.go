package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	if c.Session["email"] == "" {
		return c.Redirect(Accounts.Login)
	}

	return c.Render()
}
