import QtQuick 2.7          //Rectangle
import QtQuick.Controls 2.1 //StackView

import Theme 1.0            //Theme
import ViewTemplate 1.0     //StackTemplate
import Recipes 1.0        //Recipes

StackTemplate {
  Rectangle {
    anchors.fill: parent
    color: Theme.darkBackground

    StackView {
      id: stackView
      anchors.fill: parent

      initialItem: recipes

      Recipes { id: recipes }
//      Favorites { id: favorites }
    }
  }

  onClicked: {
    var next = nextItem(code)
    if (next != stackView.currentItem) {
      stackView.replace(next, StackView.Immediate)
    }
  }

  function nextItem(code) {
    switch (code) {
    case "recipes":
      return recipes
//    case "favorites":
//      return favorites
    }
  }
}
