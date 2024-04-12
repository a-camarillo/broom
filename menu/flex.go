package menu

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type menuFlex struct {
  *tview.Flex
}

func NewFlexBox(m *Menu) *menuFlex {
  menuFlex := &menuFlex{
    Flex: tview.NewFlex(),
  }
  menuFlex.setKeybinding(m)
  return menuFlex
}

func (f *menuFlex) setKeybinding(m *Menu) {
  f.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
    switch event.Key(){
    case tcell.KeyRight:
          m.app.SetFocus(f.Flex)
    } 

    switch event.Rune() {
    case '?':
      m.pages.ShowPage("help")
    case 'd':
      m.pages.ShowPage("confirmation")
    case 'h':
      m.app.SetFocus(f.GetItem(0))
    case 'l':
      m.app.SetFocus(f.GetItem(1))
    }

    return event
  }) 
}
