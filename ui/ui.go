package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

// Menu is a struct for containing the outer most parts of the visual
// menu whose fields should be accesible by children elements.
type UI struct {
  app *tview.Application
  pages *tview.Pages
}

func NewUI() *UI {
  return &UI{
    app: tview.NewApplication(),
    pages: tview.NewPages(),
  }
}

func Initialize() {
  ui := NewUI()
  
  flex := NewFlexBox(ui)
  localList := NewBranchList(ui).newLocalList()
  deleteList := NewBranchList(ui).newDeleteList()
  confirmationModal := NewMenuModal(ui)
  helpModal := NewMenuModal(ui)

  flex.AddItem(localList, 0, 1, true).
       AddItem(deleteList, 0, 1, false)

  ui.pages.
    AddPage("container", flex, true, true).
    AddPage("help",
    helpModal.NewHelpModal(ui), true, true).
    AddPage("confirmation", confirmationModal.NewConfirmationModal(ui),
    true,
    false)
  
  if err := ui.app.SetRoot(ui.pages, true).Run(); err != nil {
    errString := fmt.Sprintf("%s", err)
    panic(errString)
  } 
}
 
