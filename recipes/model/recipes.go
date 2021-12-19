package model

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/therecipe/qt/core"
)

type TCookingDish int
type TCookingMethod int

//soup hot dessert stove for cooking multicooker microwave
const (
	SOUP TCookingDish = iota + 1
	HOT
	DESSERT

	COOKER TCookingMethod = iota + 1
	STOVE
	MICROWAVE
	MULTI_COOKER
)

func (d TCookingDish) String() string {
	return [...]string{"soup", "hot", "dessert"}[d-1]
}

func (m TCookingMethod) String() string {
	return [...]string{"cooker", "stove", "microwave", "multi_cooker"}[m-1]
}

func (d TCookingDish) EnumIndex() int {
	return int(d)
}

func (m TCookingMethod) EnumIndex() int {
	return int(m)
}

type Recipe struct {
	Title          string  `json:"title"`
	Picture        string  `json:"picture"`
	ReadyInMinutes float64 `json:"readyInMinutes"`
	CookingDish    string  `json:"cookingDish"`
	CookingMethod  string  `json:"cookingMethod"`
	Rating         float64 `json:"rating"`

	ID               uuid.UUID                     `json:"id"`
	Ingredients      []string                      `json:"ingredients"`
	IngredientMetric map[string]IngredientsDetails `json:"ingredientInfo"`
	Instructions     string                        `json:"summary"`
	Steps            []Step                        `json:"steps"`
}

type Step struct {
	Image       string `json:"image"`
	Description string `json:"description"`
}

type IngredientsDetails struct {
	Amount       float64 `json:"amount"`
	MeasuresUnit string  `json:"measuresUnit"`
	Image        string  `json:"image"`
	Name         string  `json:"name"`
}

type RecipesModel struct {
	core.QAbstractTableModel

	recipes []Recipe

	Filter *core.QSortFilterProxyModel

	_ func() `constructor:"init"`

	_ func([]Recipe) `signal:"updateWith,auto"`
}

func (m *RecipesModel) init() {
	m.Filter = core.NewQSortFilterProxyModel(nil)
	m.Filter.SetSourceModel(m)
	m.Filter.SetFilterKeyColumn(0)
	m.Filter.SetFilterCaseSensitivity(core.Qt__CaseInsensitive)

	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(func(*core.QModelIndex) int { return 5 }) //needed for sort filter model
	m.ConnectData(m.data)
	m.ConnectRoleNames(m.roleNames)
}

func (m *RecipesModel) rowCount(parent *core.QModelIndex) int {
	return len(m.recipes)
}

func (m *RecipesModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) && role < int(core.Qt__UserRole) {
		return core.NewQVariant()
	}

	dbArray := m.recipes
	if index.Row() < 0 || index.Row() >= len(dbArray) {
		return core.NewQVariant()
	}

	dbItem := dbArray[index.Row()]

	switch {
	case index.Column() == 0 && role == int(core.Qt__DisplayRole) || //needed for sort filter model
		role == int(core.Qt__UserRole)+1:
		return core.NewQVariant1(dbItem.Title)

	case role == int(core.Qt__UserRole)+2:
		return core.NewQVariant1(dbItem.Picture)

	case role == int(core.Qt__UserRole)+3:
		return core.NewQVariant1(dbItem.ReadyInMinutes)

	case role == int(core.Qt__UserRole)+4:
		return core.NewQVariant1(dbItem.CookingDish)

	case role == int(core.Qt__UserRole)+5:
		return core.NewQVariant1(dbItem.CookingMethod)

	case role == int(core.Qt__UserRole)+6:
		return core.NewQVariant1(map[string]*core.QVariant{
			"text":  core.NewQVariant1(fmt.Sprintf("%.1f%", dbItem.Rating)),
			"value": core.NewQVariant1(dbItem.Rating),
		})

	case role == int(core.Qt__UserRole)+7:
		return core.NewQVariant1(map[string]*core.QVariant{
			"available": core.NewQVariant9(true),
		})

	case role == int(core.Qt__UserRole)+8:
		return core.NewQVariant1(dbItem.ID.String())
	}

	return core.NewQVariant()
}

func (m *RecipesModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		int(core.Qt__UserRole) + 1: core.NewQByteArray2("TITLE", -1),
		int(core.Qt__UserRole) + 2: core.NewQByteArray2("PICTURE", -1),
		int(core.Qt__UserRole) + 3: core.NewQByteArray2("READY_IN_MINUTES", -1),
		int(core.Qt__UserRole) + 4: core.NewQByteArray2("COOKING_DISH", -1),
		int(core.Qt__UserRole) + 5: core.NewQByteArray2("COOKING_METHOD", -1),
		int(core.Qt__UserRole) + 6: core.NewQByteArray2("RATING", -1),
		int(core.Qt__UserRole) + 7: core.NewQByteArray2("ACTIONS", -1),
	}
}

func (m *RecipesModel) updateWith(recipes []Recipe) {
	if len(m.recipes) != len(recipes) {
		m.Filter.BeginResetModel()
		m.recipes = recipes
		m.Filter.EndResetModel()
	} else {
		m.recipes = recipes
		m.Filter.DataChanged(m.Filter.Index(0, 0, core.NewQModelIndex()), m.Filter.Index(len(m.recipes)-1, 0, core.NewQModelIndex()), make([]int, 0))
	}
}
