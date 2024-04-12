package menu

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type menuModal struct {
  *tview.Modal
}

func NewMenuModal(m *Menu) *menuModal {
  return &menuModal{
    Modal: tview.NewModal(),
  } 
}

func (m *menuModal) NewConfirmationModal() *tview.Modal {
  // handle all Modal Methods
  confirmationModal := m.Modal
  confirmationModal.SetText("Are you sure you want to delete these branches?").
  SetBackgroundColor(tcell.ColorBlack).
  AddButtons([]string{"Yes", "No"})

  // handle all Box Methods
  return confirmationModal
}
