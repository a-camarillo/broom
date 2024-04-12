package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type uiModal struct {
  *tview.Modal
}

func NewMenuModal(u *UI) *uiModal {
  return &uiModal{
    Modal: tview.NewModal(),
  } 
}

func (m *uiModal) NewConfirmationModal(u *UI) *uiModal {
  // handle all Modal Methods
  confirmationModal := &uiModal{
      Modal: tview.NewModal(),
  }
  confirmationModal.SetText("Are you sure you want to delete these branches?").
  SetBackgroundColor(tcell.ColorBlack).
  AddButtons([]string{"Yes", "No"})

  // handle all Box Methods
  return confirmationModal
}

func (m *uiModal) NewHelpModal(u *UI) *uiModal {
  helpBox := &uiModal{
    Modal: tview.NewModal(),
  }
  helpBox.setKeybinding(u)
  helpBox.SetBackgroundColor(tcell.Color(tcell.ColorBlack))
  helpBox.SetText(`
  To quit use "Ctrl+c" or "q"
  To toggle this help box, press "?"
  To navigate between "Current Branches" and "Branches To Be Deleted" use h/◀ and l/▶
  To navigate the individual branch lists use j/ and k/
  To add/remove a branch from "Branches To Be Deleted", highlight the current branch and press Enter
  `)
  

  return helpBox
}


func (m *uiModal) confirmationDoneFunc(u *UI) {
  m.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
    if buttonLabel == "No" {
      u.app.Stop()
    } else {
      u.app.Stop()
    }
  })
}

func (m *uiModal) setKeybinding(u *UI) {
  m.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
    switch event.Rune() {
    case '?':
      u.pages.HidePage("help")
    }

    return event 
  })
}
