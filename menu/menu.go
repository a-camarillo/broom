package menu

import (
	"fmt"

	"github.com/rivo/tview"
)

// Menu is a struct for containing the outer most parts of the visual
// menu whose fields should be accesible by children elements.
type Menu struct {
  app *tview.Application
  pages *tview.Pages
}

func NewMenu() *Menu {
  return &Menu{
    app: tview.NewApplication(),
    pages: tview.NewPages(),
  }
}

func Initialize() {
  menu := NewMenu()
  
  flex := NewFlexBox(menu)
  helpBox := NewHelpBox(menu)
  localList := NewBranchList(menu).newLocalList()
  deleteList := NewBranchList(menu).newDeleteList()
  confirmationModal := NewMenuModal(menu)

  flex.AddItem(localList, 0, 1, true).
       AddItem(deleteList, 0, 1, false)

  menu.pages.
    AddPage("container", flex, true, true).
    AddPage("help", helpBox, true, true).
    AddPage("confirmation", confirmationModal.NewConfirmationModal(),
    true,
    false)
  
  if err := menu.app.SetRoot(menu.pages, true).Run(); err != nil {
    errString := fmt.Sprintf("%s", err)
    panic(errString)
  } 
}
 
