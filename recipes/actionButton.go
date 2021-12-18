package recipes

import (
	"github.com/therecipe/qt/quick"

	"github.com/lol1pop/recipes_book/recipes/controller"
)

func init() { actionButtonTemplate_QmlRegisterType2("RecipesTemplate", 1, 0, "ActionButtonTemplate") }

type actionButtonTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(string) `signal:"showDownload,->(controller.ActionButtonController)"`
	_ func(string) `signal:"deleteRequest,->(controller.ActionButtonController)"`
}

func (t *actionButtonTemplate) init() {
	if controller.ActionButtonController == nil {
		controller.NewActionButtonController(nil)
	}
}
