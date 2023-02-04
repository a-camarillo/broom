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

func Init() {
	app := tview.NewApplication()

	branches := GetBranches()

	helpBox := tview.NewTextView()
	helpBox.SetBorder(true).SetTitle("Help")
	helpBox.SetText(`
	To Add/Remove a branch from "Branches To Be Deleted", highlight the current branch and press Enter
	`)

	branchList := tview.NewList().ShowSecondaryText(false)
	branchList.SetBorder(true).SetTitle("Current Branches")

	deleteList := tview.NewList().ShowSecondaryText(false)
	deleteList.SetSelectedFocusOnly(true).SetMainTextColor(tcell.ColorRed.TrueColor())
	deleteList.SetBorder(true).SetTitle("Branches To Be Deleted")

	branchList.SetSelectedFunc(func(idx int, main string, secondary string, shortcut rune) {
		if deleteList.FindItems(main, "", false, false) == nil {
			deleteList.AddItem(main, "", 0, nil)
		} else {
			deleteList.RemoveItem(idx)
		}
	})

	for i := range branches {
		branchList.AddItem(branches[i], "", 0, nil)
	}

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(branchList, 0, 1, true).
			AddItem(deleteList, 0, 1, false), 0, 3, true).
		AddItem(helpBox, 0, 1, false)
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
