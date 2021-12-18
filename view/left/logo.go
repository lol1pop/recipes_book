package left

import (
	"github.com/therecipe/qt/quick"

	_ "github.com/lol1pop/recipes_book/view/left/controller"
)

func init() { logoTemplate_QmlRegisterType2("LeftTemplate", 1, 0, "LogoTemplate") }

type logoTemplate struct {
	quick.QQuickItem

	_ func() `signal:"clicked,->(controller.NewLogoController(nil))"`
}
