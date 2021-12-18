package view

import (
	"github.com/therecipe/qt/quick"

	_ "github.com/lol1pop/recipes_book/view/controller"
)

func init() { viewTemplate_QmlRegisterType2("ViewTemplate", 1, 0, "ViewTemplate") }

type viewTemplate struct {
	quick.QQuickItem

	_ func(b bool) `signal:"Blur,<-(controller.NewViewController(nil))"`
}
