package menu

import (
 "github.com/rivo/tview"
)

type helpBox struct {
  *tview.TextView
}

func NewHelpBox(m *Menu) *helpBox {
  helpBox := &helpBox{
    TextView: tview.NewTextView(),
  }
  helpBox.SetText(`
  To quit use "Ctrl+c" or "q"
  To toggle this help box, press "?"
  To navigate between "Current Branches" and "Branches To Be Deleted" use h/◀ and l/▶
  To navigate the individual branch lists use j/ and k/
  To add/remove a branch from "Branches To Be Deleted", highlight the current branch and press Enter
  `)
  

  return helpBox
}
