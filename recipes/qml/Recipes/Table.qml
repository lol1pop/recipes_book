import QtQuick.Controls 1.4         //TableView
import QtQuick.Controls.Styles 1.4  //TableViewStyle

import Theme 1.0                     //Theme
import RecipesTemplate 1.0

TableView {
  id: tableView

  property RecipesTemplate template

  backgroundVisible: false
  frameVisible: false

  horizontalScrollBarPolicy: Qt.ScrollBarAlwaysOff
  verticalScrollBarPolicy: Qt.ScrollBarAlwaysOff

  model: RecipesModel

  style: TableViewStyle {

    headerDelegate: TableHeaderDelegate {}

    rowDelegate: TableRowDelegate {}
  }

  TableColumnDelegate { role: "TITLE" }

  TableColumnDelegate { role: "PICTURE" }

  TableColumnDelegate { role: "READY_IN_MINUTES" }

  TableColumnDelegate { role: "COOKING_DISH" }

  TableColumnDelegate { role: "COOKING_METHOD" }

  TableColumnDelegate { role: "RATING" }

  TableColumnActionDelegate {
    role: "ACTIONS"
    tableView: tableView
  }
}
