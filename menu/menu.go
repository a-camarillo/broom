package menu

import (
  "github.com/rivo/tview"
)

// Menu is a struct for containing the outer most parts of the visual
// menu whose fields should be accesible by children elements.
type Menu struct {
  App *tview.Application
  pages *tview.Pages
}

func NewMenu() *Menu {
  return &Menu{
    App: tview.NewApplication(),
  }
}
 
