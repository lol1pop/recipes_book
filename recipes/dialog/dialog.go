package dialog

import (
	"github.com/therecipe/qt/core"

	fcontroller "github.com/lol1pop/recipes_book/recipes/dialog/controller"
)

type filesDialogTemplate struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(cident string) `signal:"show,<-(fcontroller.Controller)"`
	_ func(bool)          `signal:"blur,->(fcontroller.Controller)"`
}

func (t *filesDialogTemplate) init() {
	if fcontroller.Controller == nil {
		fcontroller.NewDialogController(nil)
	}
}

func (t *filesDialogTemplate) show(cident string) {
	if fcontroller.Controller.IsLocked() {
	} else {
		t.Show(cident)
	}
}
