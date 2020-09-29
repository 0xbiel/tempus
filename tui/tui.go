package tui

import (
  "log"
  ui "github.com/gizak/termui/v3"
  "github.com/gizak/termui/v3/widgets"
)

func Start() {
  if err := ui.Init(); err != nil {
    log.Fatalf("failed to initialize termui: %v", err) 
  } 
  defer ui.Close()

  p := widgets.NewParagraph()
	p.Title = "Tempus"
	p.Text = "by 0xbiel."
	p.SetRect(0, 0, 15, 3)

	listData := []string{
		"Studying",
		"Working",
		"Hacking",
		"Programming",
		"Reading",
		"Writing",
	}

	l := widgets.NewList()
  l.Title = "Tasks"
  l.Rows = listData
  l.SetRect(0, 3, 25, 13)

  p2 := widgets.NewParagraph()
  p2.Title = "Active Task"
  p2.Text = "Programming"
  p2.SetRect(0, 13, 25, 17)

  ui.Render(p, l, p2)

  for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
