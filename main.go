package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

type Window struct {
	*walk.MainWindow
}

var (
	// Global controls
	DeskResolutionX int
	DeskResolutionY int
)

func main() {
	var err error
	w := new(Window)
	DeskResolutionX = int(win.GetSystemMetrics(win.SM_CXSCREEN))
	DeskResolutionY = int(win.GetSystemMetrics(win.SM_CYSCREEN))

	_, err = MainWindow{
		AssignTo: &w.MainWindow,
		Title:    "IrcGoClient",
		MinSize:  Size{Width: DeskResolutionX / 4, Height: DeskResolutionY / 4},
		Layout:   VBox{},
		Children: []Widget{
			VSplitter{Children: []Widget{
				TextEdit{
					VScroll: true,
				},
				LineEdit{},
			}},
		},
	}.Run()
	if err != nil {
		fmt.Printf("MainWindow ERROR: %s", err)
	}
}
