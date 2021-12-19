import QtQuick 2.7           //Item
import RecipesTemplate 1.0   //RecipesTemplate
import Dialog 1.0            //RecipeInfo

RecipesTemplate {
  id: template

  Info {}

  Header {
    id: header
  }

  Item {
    anchors {
      top: header.bottom
      left: parent.left
      right: parent.right
      bottom: parent.bottom
      leftMargin: 5
      rightMargin: 5
    }

    Table {
      anchors.fill: parent
      template: template
    }
  }
}
