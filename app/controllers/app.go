package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (controller App) Index() revel.Result {
	return controller.Render()
}

func (controller App) About() revel.Result {
	return controller.Render()
}

func (controller App) Contact() revel.Result {
	return controller.Render()
}

func (controller App) Blog() revel.Result {
	return controller.Render()
}
