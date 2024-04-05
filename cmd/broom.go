package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeMenu() {

	var branchesToDelete []string
	var path string
        
        pwd, _ := os.Getwd()

	modalIsOpen := false

	app := tview.NewApplication()

//	if len(os.Args) > 1 {
//		path = os.Args[1]
//	} else {
//		path = "."
//	}
        path = pwd
	repo, err := NewGitRepositoryFromString(path)
	if err != nil {
          error := fmt.Errorf("error for path %s: %s", pwd, err)
          log.Fatal(error)
	}

	refs, err := NewReferences(repo)
	if err != nil {
          error := fmt.Errorf("error: %s", err)
          log.Fatal(error)
	}

	helpBox := tview.NewTextView()
	helpBox.SetBorder(true).SetTitle("Help")
	helpBox.SetText(`
	To quit use Ctrl + C
	To toggle this help box, press "?"
	To navigate between "Current Branches" and "Branches To Be Deleted" use h/◀ and l/▶
	To add/remove a branch from "Branches To Be Deleted", highlight the current branch and press Enter
	To confirm branches to be deleted, press "/" 
	`)

	branchList := tview.NewList().ShowSecondaryText(false)
	branchList.SetBorder(true).SetTitle("Current Local Branches")

	deleteList := tview.NewList().ShowSecondaryText(false)
	deleteList.SetSelectedFocusOnly(true).SetMainTextColor(tcell.ColorRed.TrueColor())
	deleteList.SetBorder(true).SetTitle("Branches To Be Deleted")

	branchList.SetSelectedFunc(func(idx int, main string, sec string, short rune) {
		if deleteList.FindItems(main, "", false, false) == nil {
			deleteList.AddItem(main, "", 0, nil)
			branchesToDelete = append(branchesToDelete, main)
		}
	})

	deleteList.SetSelectedFunc(func(idx int, main string, sec string, short rune) {
		deleteList.RemoveItem(idx)
		branchesToDelete = pop(branchesToDelete, idx)
	})

	modal := tview.NewModal().
		SetText("Are you sure you want to delete these branches?").
		SetBackgroundColor(tcell.ColorBlack).
		AddButtons([]string{"Yes", "No"})

	refNames, _ := refs.GetReferenceNames()
	for _, i := range refNames {
		branchList.AddItem(i, "", 0, nil)
	}

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(branchList, 0, 1, true).
			AddItem(deleteList, 0, 1, false), 0, 3, true).
		AddItem(helpBox, 0, 1, false)

	pages := tview.NewPages().
		AddPage("interface", flex, true, true).
		AddPage("modal", modal, true, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'l' {
			app.SetFocus(deleteList)
			return tcell.NewEventKey(tcell.KeyRune, 'l', tcell.ModNone)
		}
		if event.Key() == tcell.KeyRight {
			app.SetFocus(deleteList)
			return tcell.NewEventKey(tcell.KeyRune, rune(tcell.KeyRight), tcell.ModNone)
		}
		if event.Rune() == 'h' {
			app.SetFocus(branchList)
			return tcell.NewEventKey(tcell.KeyRune, 'h', tcell.ModNone)
		}
		if event.Key() == tcell.KeyLeft {
			app.SetFocus(branchList)
			return tcell.NewEventKey(tcell.KeyRune, rune(tcell.KeyLeft), tcell.ModNone)
		}
		if event.Rune() == '?' {
			if flex.GetItemCount() > 1 {
				flex.RemoveItem(helpBox)
				return tcell.NewEventKey(tcell.KeyRune, '?', tcell.ModNone)
			} else {
				flex.AddItem(helpBox, 0, 1, false)
			}
		}
		if event.Rune() == '/' {
			if !modalIsOpen {
				pages.ShowPage("modal")
				modalIsOpen = true
			}
		}
		return event
	})

	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "No" {
			pages.HidePage("modal")
			modalIsOpen = false
		} else {
			for _, i := range branchesToDelete {
				repo.Repository.DeleteBranch(i)
			}
			app.Stop()
		}
	})

	if err := app.SetRoot(pages, true).Run(); err != nil {
		errString := fmt.Sprintf("%s", err)
		panic(errString)
	}
}
