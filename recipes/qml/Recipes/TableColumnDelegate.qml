import QtQuick 2.7          //Component
import QtQuick.Controls 1.4 //TableViewColumn
import QtQuick.Controls 2.1 //Label

import Theme 1.0            //Theme

TableViewColumn {

  title: role
  resizable: true
  movable: false
  width: parent.width * 0.25

  delegate: Component {
      Item {
        id: cell
          Label {
            anchors.centerIn: parent

            color: styleData.selected ? Theme.walletTableHighlight : styleData.column == 0 || styleData.column == 7 ? Theme.font : Theme.fontHighlight
            text: styleData.value
            font.bold: styleData.column == 0
            font.italic: styleData.column == 3

            visible: styleData.column != 1
          }

          Image {
            source: styleData.value
            cache: false

            fillMode: Image.PreserveAspectFit;
            width: cell.width;
            height: cell.height

            visible: styleData.column == 1
          }
      }
  }
}
