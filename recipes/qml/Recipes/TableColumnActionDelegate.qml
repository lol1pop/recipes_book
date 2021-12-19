import QtQuick 2.7          //Component
import QtQuick.Controls 1.4 //TableViewColumn
import QtQuick.Controls 2.1 //Label

import Theme 1.0            //Theme
import RecipesTemplate 1.0

import "."                  //ProgressBar

TableViewColumn {

  property QtObject tableView
  property ActionButtonTemplate template: ActionButtonTemplate{}

  title: role
  resizable: false
  movable: false
  width: parent.width * 0.25

  delegate: Component { Item {

        ProgressBar {
          id: progressBar

          anchors {
            verticalCenter: parent.verticalCenter
            left: parent.left
          }
          width: parent.width * 0.75
          height: parent.height * 0.6

          progressBarText: styleData.value.text != null ? styleData.value.text : ""
          value: styleData.value.value != null ? styleData.value.value : 0

          visible: !downloadButton.visible || !deleteButton.visible
        }

        TableColumnDelegateButton {
          id: downloadButton

          anchors {
            top: parent.top
            left: parent.left
            bottom: parent.bottom
            }
          width: parent.width * 0.75

          source: "qrc:/qml/assets/ic_cloud_download_black_24px.svg"
          accent: Theme.accent
          onClicked: template.showDownload(tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1))

          visible: styleData.value.available != null ? styleData.value.available : false
        }

        TableColumnDelegateButton {
          id: deleteButton
          anchors {
            top: parent.top
            left: downloadButton.right
            right: parent.right
            bottom: parent.bottom
          }

          source: "qrc:/qml/assets/ic_delete_forever_black_24px.svg"
          accent: Theme.walletTableHighlight
          onClicked: template.deleteRequest(tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1))

          visible: styleData.value.available != null ? styleData.value.available : false
        }
    }
  }
}
