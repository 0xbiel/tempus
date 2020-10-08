package tui

import (
//  "fmt"
  "log"
  "github.com/jroimartin/gocui"
)

func Start() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout) //set layout manager that is the layout function.

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

  mainView, err := g.SetView("Main", 0, 0, maxX, maxY - 4)

  if (err != nil && err != gocui.ErrUnknownView) {
		log.Println("Failed to create main view:", err)
		return
	}

  mainView.Title = "Main"
  mainView.FgColor = gocui.ColorWhite

  v, _ := g.SetCurrentView(mainView)
  //cursorUp(g, v)

 // err = g.MainLoop()
	log.Println("Main loop has finished:", err)

    //return
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
