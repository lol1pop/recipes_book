package view

import (
	"github.com/therecipe/qt/quick"

	_ "github.com/lol1pop/recipes_book/view/controller"
)

func init() { stackTemplate_QmlRegisterType2("ViewTemplate", 1, 0, "StackTemplate") }

type stackTemplate struct {
	quick.QQuickItem

	_ func(code string) `signal:"Clicked,<-(controller.NewStackController(nil))"`
}
