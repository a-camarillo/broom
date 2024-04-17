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
  confirmationModal.setKeybinding(u)
  confirmationModal.confirmationDoneFunc(u)
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
      u.pages.HidePage("confirmation")
    } else {
      deleteBranches(u) 
      u.app.Stop()
    }
  })
}

func (m *uiModal) setKeybinding(u *UI) {
  m.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
    switch event.Rune() {
    case 'h':
      return tcell.NewEventKey(tcell.KeyTab, 0, tcell.ModNone) 
    case 'l':
      return tcell.NewEventKey(tcell.KeyBacktab, 0, tcell.ModNone) 
    case '?':
      u.pages.HidePage("help")
    case ' ':
      return tcell.NewEventKey(tcell.KeyEnter, 0, tcell.ModNone)
    }

    return event 
  })
}
