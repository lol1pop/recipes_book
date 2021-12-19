import QtQuick 2.7  //Item

import Theme 1.0    //Theme
import View 1.0     //StackItemHeader

StackItemHeader {
  image: "qrc:/qml/assets/ic_insert_drive_file_black_24px.svg"
  title: "Recipes"

   contentItem: Item {
    anchors.fill: parent

    HeaderButton {
      anchors {
        top: parent.top
        right: create.left
        bottom: parent.bottom
      }
      width: Theme.minimumWidth * 0.1

      text: "Info"
      code: "info"
      image: "qrc:/qml/assets/ic_share_black_24px.svg"
    }

    HeaderButton {
      id: create
      anchors {
        top: parent.top
        right: parent.left
        bottom: parent.bottom
      }
      width: Theme.minimumWidth * 0.1

      text: "Create Recipes"
      code: "create"
      image: "qrc:/qml/assets/ic_create_black_24px.svg"
    }
  }
}
