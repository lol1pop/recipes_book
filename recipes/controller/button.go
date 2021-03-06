package controller

import (
	"github.com/therecipe/qt/core"

	_ "github.com/lol1pop/recipes_book/recipes/dialog/controller"
)

var ButtonController *buttonController

type buttonController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked,->(controller.Controller.Show)"`
}

func (c *buttonController) init() {
	ButtonController = c
}
