import QtQuick 2.7              //TextInput
import QtQuick.Layouts 1.3      //Layout

import Theme 1.0                //Theme
import DialogTemplate 1.0       //ReceiveTemplate

import "."                      //Dialog

Dialog {
    id: dialog
    property InfoTemplate template: InfoTemplate {  onShow: dialog.visible = (cident == "info"); }
      customTitle: "You can receive Coins using the following address:"
      customLabel: "Enter a seed to recover funds from."

      customContentItem: Input {
        id: textField

        Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
        Layout.preferredWidth: parent.width * 0.8

//        text: template.address
//
//        readOnly: false
//        selectionColor: Theme.font
//        selectByMouse: true

        onAccepted: dialog.accept()
      }

        onAccepted: {
        console.log("askjhgdfaghajklj;akjbhvdgxhfxgchjk")
          var ret = template.info(textField.text)
          if (!ret[0]) {
            dialog.visible = true

            dialog.customError = ret[1]
            textField.clear()
          } else {
            dialog.reject()
          }
        }

        onRejected: {
          dialog.customError = ""
          textField.clear()
        }

    onVisibleChanged: template.blur(visible)
}