package app

import (
	"log"
	"os/exec"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GetBranches() []string {
	output, err := exec.Command("git", "branch").Output()
	if err != nil {
		log.Fatal(err)
	}
	branchStr := strings.TrimSpace(string(output))
	branchSlice := strings.Split(branchStr, "\n")
	var trimSlice []string
	for i := range branchSlice {
		trimSlice = append(trimSlice, strings.TrimSpace(branchSlice[i]))
	}
	return trimSlice
}

func pop(s []string, i int) []string {
	s[i] = s[len(s)-1]
	s[len(s)-1] = ""
	s = s[:len(s)-1]
	return s
}

func Init() {

	var branchesToDelete []string
	modalIsOpen := false

	app := tview.NewApplication()

	branches := GetBranches()

	helpBox := tview.NewTextView()
	helpBox.SetBorder(true).SetTitle("Help")
	helpBox.SetText(`
	To quit use Ctrl + C
	To toggle this help box, press "?"
	To navigate between "Current Branches" and "Branches To Be Deleted" use h/◀ and l/▶
	To add/remove a branch from "Branches To Be Deleted", highlight the current branch and press Enter 
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

	for i := range branches {
		branchList.AddItem(branches[i], "", 0, nil)
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
			deleteString := strings.Join(branchesToDelete, " ")
			err := exec.Command("git", "branch", "-d", deleteString).Run()
			if err != nil {
				panic(err)
			}
			app.Stop()
		}
	})

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
