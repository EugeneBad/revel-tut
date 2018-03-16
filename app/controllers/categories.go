package controllers

import "github.com/revel/revel"

type Categories struct {
	*revel.Controller
}

func (c Categories) Categories() revel.Result {
	return c.Render()
}