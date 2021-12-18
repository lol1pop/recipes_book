package recipes

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"

	"github.com/lol1pop/recipes_book/recipes/controller"
	"github.com/lol1pop/recipes_book/recipes/dialog"
)

func init() { recipesTemplate_QmlRegisterType2("RecipesTemplate", 1, 0, "RecipesTemplate") }

type recipesTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"RecipesModel"`
}

func (t *recipesTemplate) init() {
	c := controller.NewRecipesController(nil)

	t.SetRecipesModel(c.Model().Filter)

	//needed here, because those are non qml views
	dialog.NewFilesUploadTemplate(nil)
	dialog.NewFolderUploadTemplate(nil)
}
