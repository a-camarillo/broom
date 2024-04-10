package menu

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type branchList struct {
  *tview.List
}

func NewBranchList(m *Menu) *branchList {
  branchList := &branchList{
    List: tview.NewList().ShowSecondaryText(false),
  }
  branchList.setKeybinding(m)
  return branchList
}

func (b *branchList) newLocalList() *tview.List {
  // handle all List Methods
  localList := tview.NewList().
  ShowSecondaryText(false)

  // handle all Box Methods
  localList.SetBorder(true).
  SetTitle("Current Local Branches")

  return localList
}

func (b *branchList) newDeleteList() *tview.List {
  // handle all List Methods
  deleteList := b.List.
  ShowSecondaryText(false).
  SetSelectedFocusOnly(true).
  SetMainTextColor(tcell.ColorRed.TrueColor())


  // handle all Box Methods
  deleteList.SetBorder(true).
  SetTitle("Branches To Be Deleted")

  return deleteList
}

func (b *branchList) setKeybinding(m *Menu) {
  b.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
    switch event.Rune() {
    case 'j':
      b.moveDown()
    case 'k':
      b.moveUp()
    }
    return event
  })
}

func (b *branchList) moveDown() {
  length := b.GetItemCount() - 1

  if b.GetCurrentItem() != length {
    b.SetCurrentItem(b.GetCurrentItem()+1)
  } else {
    b.SetCurrentItem(0)
  }
}

func (b *branchList) moveUp() { 
  length := b.GetItemCount() - 1
  if b.GetCurrentItem() != 0 { 
    b.SetCurrentItem(b.GetCurrentItem()-1)
  } else {
    b.SetCurrentItem(length)
  }
}
