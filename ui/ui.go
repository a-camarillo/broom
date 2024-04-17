package ui

import (
	"fmt"
	"log"
	"os"

	"github.com/a-camarillo/broom/branch"
	"github.com/rivo/tview"
)

// Menu is a struct for containing the outer most parts of the visual
// menu whose fields should be accesible by children elements.
type UI struct {
  app *tview.Application
  pages *tview.Pages
  repo *branch.GitRepository
}

func NewUI() *UI {
  return &UI{
    app: tview.NewApplication(),
    pages: tview.NewPages(),
    repo: getRepo(),
  }
}

func Initialize(remotes bool) {
  ui := NewUI()
  
  flex := NewFlexBox(ui)
  branchList := NewBranchList(ui, remotes)
  confirmationModal := NewMenuModal(ui)
  helpModal := NewMenuModal(ui)

  flex.AddItem(branchList.l.List, 0, 1, true).
       AddItem(branchList.d.List, 0, 1, true)

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

func getRepo() *branch.GitRepository {
  pwd, _ := os.Getwd()
  
  repo, err := branch.NewGitRepositoryFromString(pwd)
  if err != nil {
    error := fmt.Errorf("error for path %s: %s", pwd, err)
    log.Fatal(error)
  }
  return repo
}
