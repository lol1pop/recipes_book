package recipes

import (
	"github.com/therecipe/qt/quick"

	"github.com/lol1pop/recipes_book/recipes/controller"
)

func init() { buttonTemplate_QmlRegisterType2("RecipesTemplate", 1, 0, "ButtonTemplate") }

type buttonTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked,->(controller.ButtonController)"`
}

func (t *buttonTemplate) init() {
	if controller.ButtonController == nil {
		controller.NewButtonController(nil)
	}
}
