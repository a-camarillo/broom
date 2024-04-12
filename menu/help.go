package menu

import (
 "github.com/rivo/tview"
 "github.com/gdamore/tcell/v2"
)

type helpBox struct {
  *tview.Modal
}

func NewHelpBox(m *Menu) *helpBox {
  helpBox := &helpBox{
    Modal: tview.NewModal(),
  }
  helpBox.setKeybinding(m)
  helpBox.SetText(`
  To quit use "Ctrl+c" or "q"
  To toggle this help box, press "?"
  To navigate between "Current Branches" and "Branches To Be Deleted" use h/◀ and l/▶
  To navigate the individual branch lists use j/ and k/
  To add/remove a branch from "Branches To Be Deleted", highlight the current branch and press Enter
  `)
  

  return helpBox
}

func (h *helpBox) setKeybinding(m *Menu) {
  h.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
    switch event.Rune() {
    case '?':
      m.pages.HidePage("help")
    }

    return event 
  })
}
