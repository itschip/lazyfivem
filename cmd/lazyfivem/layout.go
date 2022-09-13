package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {

	maxX, maxY := g.Size()
	if s, err := g.SetView("side", 0, 0, maxX/4, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		s.Title = "Servers"
		s.FgColor = gocui.ColorWhite
		s.SelFgColor = gocui.ColorGreen
		s.Highlight = true

    for key := range Servers {
      fmt.Fprintln(s, key)
    }

		if _, err := g.SetCurrentView("side"); err != nil {
			return err
		}
	}

	if v, err := g.SetView("main", int(0.2*float32(maxX)), 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "FXServer"
		v.Autoscroll = true
		v.Wrap = true

		fmt.Fprintln(v, "")
	}

	if c, err := g.SetView("command", int(0.2*float32(maxX)), maxY-4, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		c.Title = "Command"
		c.Editable = true
		g.Cursor = true
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
