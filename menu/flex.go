package menu

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type menuFlex struct {
  *tview.Flex
}

func NewFlexBox(m *Menu) *menuFlex {
  return &menuFlex{
    Flex: tview.NewFlex(),
  }
}

func (f *menuFlex) setKeybinding(m *Menu) {
  f.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
    return event
  }) 
}
