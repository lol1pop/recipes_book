package controller

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"

	"github.com/lol1pop/recipes_book/recipes/model"
)

var RController *RecipesController

type RecipesController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ *model.RecipesModel `property:"model"`

	_ func(ID string) `signal:"doubleClicked,auto"`
}

func (c *RecipesController) init() {
	RController = c

	c.SetModel(model.NewRecipesModel(nil))

	go c.loop()
}

func (c *RecipesController) loop() {
	//for range time.NewTicker(1 * time.Second).C {
	var df []model.Recipe
	json.Unmarshal([]byte(DEMO_RECIPES), &df)
	for i, _ := range df {
		df[i].ID = uuid.New()
	}
	c.Model().UpdateWith(df)
	fmt.Printf("loop")
	//}
}

func (c *RecipesController) doubleClicked(ID string) {
	gui.QDesktopServices_OpenUrl(core.NewQUrl3("https://example.com/"+ID, 0))
}
