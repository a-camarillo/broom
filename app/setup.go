package app

import (
	"log"
	"os/exec"
	"strings"

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

	branchList := tview.NewList().ShowSecondaryText(false)
	branchList.SetBorder(true).SetTitle("Branches")
	for i := range branches {
		branchList.AddItem(branches[i], "", 0, nil)
	}

	flex := tview.NewFlex().
		AddItem(branchList, 0, 1, true)
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
