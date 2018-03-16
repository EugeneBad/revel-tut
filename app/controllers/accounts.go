package controllers

import (
	"github.com/revel/revel"
	"github.com/eugenebad/massape/app/models"
	"fmt"
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
		fmt.Println(account)

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
		//fmt.Println(account)
		a.Validation.Required(account.Email)
		a.Validation.Required(account.Password)
		a.Validation.Email(account.Email)
		if a.Validation.HasErrors(){
			fmt.Println(a.Validation.Errors)
			return a.Render()
		}
		err := app.Db.Create(&models.Account{Email: account.Email, Password: account.Password}).GetErrors()
		if len(err) != 0{
			fmt.Println(err)
			return a.Render()
		}

		return a.RenderTemplate("App/index.html")
	}
	return nil
}