package ui

import (
	"fmt"
	"log"
	"os"

	"github.com/a-camarillo/broom/branch"
	"github.com/a-camarillo/broom/utils"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

//TODO
//deleteBranches is returning a 'branch not found' error
//need to investigate the issues but thankfully go logging
//was relatively simple to discover the error

var branchesToDelete []string

type branchList struct {
  l *localList
  d *deleteList
}

type localList struct {
  *tview.List
}

type deleteList struct {
  *tview.List
}

func NewBranchList(u *UI, remotes bool) *branchList {
  branchList := &branchList{
    l: NewLocalList(),
    d: NewDeleteList(),
  }
  branchList.setKeybinding(u)
  fillLocalList(branchList.l.List, remotes, u) 
  return branchList
}

func NewLocalList() *localList {
  // handle all List Methods
  localList := & localList{
    List: tview.NewList(),
  }

  localList.ShowSecondaryText(false)

  // handle all Box Methods
  localList.SetBorder(true).
  SetTitle("Current Local Branches")

  return localList
}

func NewDeleteList() *deleteList {
  // handle all List Methods
  deleteList := & deleteList{
    List: tview.NewList(),
  }
  deleteList.SetSelectedFocusOnly(true).
  ShowSecondaryText(false).
  SetMainTextColor(tcell.ColorRed.TrueColor())


  // handle all Box Methods
  deleteList.SetBorder(true).
  SetTitle("Branches To Be Deleted")

  return deleteList
}

func (b *branchList) setKeybinding(u *UI) {
  b.l.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
    switch event.Rune() {
    case 'j':
      moveDown(b.l.List)
    case 'k':
      moveUp(b.l.List)
    case ' ':
      b.addToDeleteList()
    }
    return event
  })
  b.d.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
    switch event.Rune() {
    case 'j':
      moveDown(b.d.List)
    case 'k':
      moveUp(b.d.List)
    case ' ':
      b.removeFromDeleteList()
    }
    return event
  })
}

func (b *branchList) addToDeleteList() {
  b.l.SetSelectedFunc(func(idx int, main string, sec string, sh rune) {
    if b.d.FindItems(main, "", false, false) == nil {
      b.d.AddItem(main, sec, sh, nil)
      branchesToDelete = append(branchesToDelete, main) 
    }
  }) 
}

func (b *branchList) removeFromDeleteList() {
  b.d.SetSelectedFunc(func(idx int, main string, sec string, sh rune) {
    b.d.RemoveItem(idx)
    branchesToDelete = utils.Pop(branchesToDelete, idx)
  })
}

func moveDown(l *tview.List) {
  length := l.GetItemCount() - 1

  if l.GetCurrentItem() != length {
    l.SetCurrentItem(l.GetCurrentItem()+1)
  } else {
    l.SetCurrentItem(0)
  }
}

func moveUp(l *tview.List) { 
  length := l.GetItemCount() - 1
  if l.GetCurrentItem() != 0 { 
    l.SetCurrentItem(l.GetCurrentItem()-1)
  } else {
    l.SetCurrentItem(length)
  }
}

func fillLocalList(l *tview.List, remotes bool, u *UI) {
  refs := initializeRefs(u)
  if !remotes {
    refNames, _ := refs.GetReferenceNames()
    for _, i := range refNames {
      l.AddItem(i, "", 0, nil)
    }
  } else {
    refNames, _ := refs.GetReferenceNamesWithRemotes()
    for _, i := range refNames {
      l.AddItem(i, "", 0, nil)
    }
  }
}

func deleteBranches(u *UI) {
  for _, i := range branchesToDelete {
    f, _ := os.OpenFile(".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    logger := log.New(f, "branches", log.LstdFlags)
    logger.Println(i)
    err := u.repo.Repository.DeleteBranch(i)
    if err != nil {
      logger.Println(err)
    }
  }
}

func initializeRefs(u *UI) *branch.References {
  refs, err := branch.NewReferences(u.repo)
  if err != nil {
    error := fmt.Errorf("error: %s", err)
    log.Fatal(error)
  }

  return refs
}
