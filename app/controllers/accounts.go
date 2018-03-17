package controllers

import (
	"github.com/revel/revel"
	"github.com/eugenebad/massape/app/models"
	"github.com/eugenebad/massape/app"
)

type Accounts struct{
	*revel.Controller
}

func (a Accounts) Login() revel.Result{
	method := a.Request.Method
	switch method {
	case "GET":
		return a.Render()
	case "POST":
		var account models.Account
		a.Params.Bind(&account, "account")

		a.Validation.Required(account.Email)
		a.Validation.Email(account.Email)
		a.Validation.Required(account.Password)

		if a.Validation.HasErrors() {
			return a.RenderText("Check email and password")
		}

		app.Db.Where("email = ? and password = ?", account.Email, account.Password).First(&account)
		if account.HasID() {
			a.Session["email"] = account.Email
			return a.RenderText("Done")
		}

		return a.RenderText("Check Email and Password")
	}
	return nil
}

func (a Accounts) Register() revel.Result{
	method := a.Request.Method
	switch method {

	case "GET":
		return a.Render()

	case "POST":
		var account models.Account
		a.Params.Bind(&account, "account")

		if a.Params.Get("confirmPassword") != account.Password{
			return a.RenderText("Passwords don't match")
		}

		a.Validation.Required(account.Email)
		a.Validation.Email(account.Email)
		a.Validation.Required(account.Password)

		if a.Validation.HasErrors(){
			return a.RenderText("Both Email and Password required")
		}
		var duplicateUser []models.Account

		if app.Db.Where("email = ?", account.Email).Find(&duplicateUser); len(duplicateUser) != 0 {
			return a.RenderText("Username Already exists")
		}

		app.Db.Create(&account)
		a.Session["email"] = account.Email
		return a.RenderTemplate("App/index.html")
	}
	return nil
}