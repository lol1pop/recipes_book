package controller

import (
	"github.com/therecipe/qt/core"

	"github.com/lol1pop/recipes_book/recipes/controller"
	lcontroller "github.com/lol1pop/recipes_book/view/left/controller"
)

type searchController struct {
	core.QObject

	_ func(string) `signal:"search,auto"`
}

func (c *searchController) search(name string) {
	lcontroller.LeftController.Clicked("recipes")
	controller.RController.Model().Filter.SetFilterFixedString(name)
}
