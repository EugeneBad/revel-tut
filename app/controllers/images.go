package controllers

import (
	"github.com/revel/revel"
	"os"
)

type Images struct {
	*revel.Controller
}

func (i Images) ImageServe() revel.Result {
	f,_ := os.Open("uploaded/QTSBsVjggSKbKiB.jpeg")
	return i.RenderFile(f, revel.Inline)

}