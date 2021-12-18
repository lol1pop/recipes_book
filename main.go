package main

import (
	"os"
	"path/filepath"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"

	_ "github.com/lol1pop/recipes_book/assets"
	_ "github.com/lol1pop/recipes_book/theme"
	_ "github.com/lol1pop/recipes_book/view"
	_ "github.com/lol1pop/recipes_book/view/left"

	_ "github.com/lol1pop/recipes_book/recipes"

	"github.com/lol1pop/recipes_book/controller"
)

var (
	pathToProject = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "lol1pop", "recipes_book")

	PRODUCTION = false //set to 'true' to use qrc: instead of qml files
	DEMO       = true  //set to 'true' to use demo data for the wallet and files table
)

func main() {
	path := filepath.Join(pathToProject, "view", "qml", "View.qml")

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	widgets.NewQApplication(len(os.Args), os.Args)

	controller.NewController(nil)

	view := quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetMinimumSize(core.NewQSize2(1024, 415))
	view.Resize(core.NewQSize2(1024, 768))

	if PRODUCTION {
		view.Engine().AddImportPath("qrc:/qml/")
		view.SetSource(core.NewQUrl3("qrc:/qml/View.qml", 0))
	} else {
		view.Engine().AddImportPath(filepath.Join(pathToProject, "theme", "qml"))

		view.Engine().AddImportPath(filepath.Join(pathToProject, "view", "qml"))
		//view.Engine().AddImportPath(filepath.Join(pathToProject, "view", "top", "qml"))
		view.Engine().AddImportPath(filepath.Join(pathToProject, "view", "left", "qml"))

		view.Engine().AddImportPath(filepath.Join(pathToProject, "recipes", "qml"))

		view.SetSource(core.QUrl_FromLocalFile(path))
	}

	view.Show()

	widgets.QApplication_Exec()
}
