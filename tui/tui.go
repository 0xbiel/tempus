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
	p.Text = "Time Tracker."
	p.SetRect(0, 0, 60, 10)

  ui.Render(p)

  for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}