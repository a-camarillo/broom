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
  confirmationModal := NewMenuModal(menu).NewConfirmationModal()

  flex.AddItem(helpBox, 0, 1, true)

  menu.pages.
    AddPage("container", flex, true, true).
    AddPage("confirmationModal", confirmationModal, true, false)
  
  if err := menu.app.SetRoot(menu.pages, true).Run(); err != nil {
    errString := fmt.Sprintf("%s", err)
    panic(errString)
  } 
}
 
