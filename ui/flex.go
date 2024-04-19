package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type uiFlex struct {
	*tview.Flex
}

func NewFlexBox(u *UI) *uiFlex {
	uiFlex := &uiFlex{
		Flex: tview.NewFlex(),
	}
	uiFlex.setKeybinding(u)
	return uiFlex
}

func (f *uiFlex) setKeybinding(u *UI) {
	f.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyLeft:
			u.app.SetFocus(f.GetItem(0))
			return tcell.NewEventKey(tcell.KeyRune, rune(tcell.KeyLeft), tcell.ModNone)
		case tcell.KeyRight:
			u.app.SetFocus(f.GetItem(1))
			return tcell.NewEventKey(tcell.KeyRune, rune(tcell.KeyRight), tcell.ModNone)
		}

		switch event.Rune() {
		case '?':
			u.pages.ShowPage("help")
		case 'd':
			u.pages.ShowPage("confirmation")
		case 'h':
			u.app.SetFocus(f.GetItem(0))
		case 'l':
			u.app.SetFocus(f.GetItem(1))
		case 'q':
			u.app.Stop()
		}

		return event
	})
}
