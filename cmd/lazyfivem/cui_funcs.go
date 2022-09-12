package main

import (
	"io"
	"log"

	"github.com/jroimartin/gocui"
)

func nextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "side" {
		_, err := g.SetCurrentView("command")
		return err
	}

	_, err := g.SetCurrentView("side")
	return err
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}

	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}

	return nil
}

func onSideEnter(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	profile, err := v.Line(cy)
	if err != nil {
		return err
	}

	startFxServer(profile, g)

	return nil
}

func executeCommand(g *gocui.Gui, v *gocui.View) error {
	var err error
	v.Rewind()
	if err != nil {
		log.Fatal(err)
	}

	if v.Buffer() != "" {
		io.WriteString(Writer, v.Buffer())
	}

	v.Clear()
	v.SetCursor(0, 0)

	if err != nil {
		return err
	}

	return nil
}
