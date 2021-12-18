package controller

import (
	"encoding/json"
	"time"

	"github.com/therecipe/qt/core"

	"github.com/lol1pop/recipes_book/recipes/model"
)

var RecipesController *recipesController

type recipesController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ *model.RecipesModel `property:"model"`
}

func (c *recipesController) init() {
	RecipesController = c

	c.SetModel(model.NewRecipesModel(nil))

	go c.loop()
}

func (c *recipesController) loop() {
	for range time.NewTicker(1 * time.Second).C {
		var df []model.Recipe
		json.Unmarshal([]byte(DEMO_RECIPES), &df)
		c.Model().UpdateWith(df)
	}
}
