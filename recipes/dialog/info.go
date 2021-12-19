package dialog

import (
	"github.com/therecipe/qt/core"

	_ "github.com/lol1pop/recipes_book/recipes/dialog/controller"
)

func init() {
	infoTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "InfoTemplate")
}

type infoTemplate struct {
	dialogTemplate

	_ func(string) *core.QVariant `slot:"info,->(controller.Controller)"`
}
