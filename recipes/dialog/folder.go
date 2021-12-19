package dialog

import (
	"github.com/therecipe/qt/widgets"

	_ "github.com/lol1pop/recipes_book/recipes/dialog/controller"
)

type folderUploadTemplate struct {
	dialogTemplate

	_ func(string, string) `signal:"download,->(controller.Controller)"`
	_ func(string)         `signal:"showDownload,<-(controller.Controller)"`
}

func (t *folderUploadTemplate) showDownload(name string) {
	t.Blur(true)
	t.Download(name, widgets.QFileDialog_GetExistingDirectory(nil, "Select a folder to download "+name, "", widgets.QFileDialog__ShowDirsOnly|widgets.QFileDialog__DontResolveSymlinks))
	t.Blur(false)
}
