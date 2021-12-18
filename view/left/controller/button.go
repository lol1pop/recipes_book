package controller

import (
	"github.com/therecipe/qt/core"

	"github.com/lol1pop/recipes_book/view/controller"
)

type buttonController struct {
	core.QObject

	_ func(code string) `signal:"clicked,auto"`
}

//lazy binding to the view/stack controller
func (c *buttonController) clicked(code string) { controller.StackController.Clicked(code) }
