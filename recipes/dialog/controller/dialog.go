package controller

import (
	"fmt"
	_ "github.com/lol1pop/recipes_book/view/controller"
	"github.com/therecipe/qt/core"
)

var Controller *dialogController

type dialogController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(cident string) `signal:"show"`
	_ func(bool)          `signal:"blur,->(controller.Controller)"`

	_ func(string)         `signal:"showDownload"`
	_ func(string, string) `signal:"download,auto"`

	_ func(ID string) *core.QVariant `slot:"info,auto"`
	//_ func() string                            `slot:"receive,auto"`
	//_ func(amount, dest string) *core.QVariant `slot:"send,auto"`
	//_ func(seed string) *core.QVariant         `slot:"recover,auto"`
}

func (c *dialogController) init() {
	Controller = c
}

func (c *dialogController) info(ID string) *core.QVariant {
	fmt.Printf("ERRRRROOORRR!!! lol net  id=" + ID)
	return core.NewQVariant()
}

func (c *dialogController) download(name, path string) {
}
