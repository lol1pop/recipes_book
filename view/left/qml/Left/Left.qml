import QtQuick 2.7            //Image
import QtQuick.Controls 2.1   //ButtonGroup
import QtQuick.Layouts 1.3    //ColumnLayout

import Theme 1.0              //Theme

import LeftTemplate 1.0       //LeftTemplate
import "." as T               //needed for name clash with std Controls.Button

LeftTemplate {
  ColumnLayout {
    anchors.fill: parent

    Item {
      Layout.preferredHeight: Theme.minHeight * 0.005
    }

    Logo {
      Layout.fillWidth: true
      Layout.preferredHeight: Theme.minHeight * 0.15
    }

    Item {
      Layout.fillHeight: true
    }

    ButtonGroup {
      buttons: column.children
    }

    ColumnLayout {
      id: column
      Layout.fillWidth: true
      Layout.preferredHeight: Theme.minHeight * 0.4

      T.Button {
        id: recipes
        Layout.fillWidth: true

        text: "Recipes"
        code: "recipes"
        image: "qrc:/qml/assets/ic_dashboard_black_24px.svg"
      }

//      T.Button {
//        id: favorites
//        Layout.fillWidth: true
//
//        text: "Favorites"
//        code: "favorites"
//        image: "qrc:/qml/assets/ic_insert_drive_file_black_24px.svg"
//      }
//
    }

    Item {
      Layout.fillHeight: true
    }

//    T.ProgressBar {
//      Layout.fillWidth: true
//      Layout.preferredHeight: Theme.minHeight * 0.125
//    }
  }

  onClicked: {
    switch (cident) {
      case "recipes":
        recipes.clicked();
        break;

//      case "favorites":
//        favorites.clicked();
//        break;

    }
  }
}
