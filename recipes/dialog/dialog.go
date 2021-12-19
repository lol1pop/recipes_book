package dialog

import (
	"github.com/lol1pop/recipes_book/recipes/dialog/controller"
	"github.com/therecipe/qt/core"
)

type dialogTemplate struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(cident string) `signal:"show,<-(controller.Controller)"`
	_ func(bool)          `signal:"blur,->(controller.Controller)"`
}

func (t *dialogTemplate) init() {
	if controller.Controller == nil {
		controller.NewDialogController(nil)
	}
}

func (t *dialogTemplate) show(cident string) {
	println("ERRRRROOORRR!!! lol net  cident=" + cident)
	t.Show(cident)
}
