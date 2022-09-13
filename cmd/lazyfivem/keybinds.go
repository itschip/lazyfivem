package main

import (

	"github.com/jroimartin/gocui"
)

func keybinds(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	// Next View
	if err := g.SetKeybinding("side", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		return err
	}

	// Next View
	if err := g.SetKeybinding("command", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		return err
	}

	// Cursor up
	if err := g.SetKeybinding("side", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}

	// Cursor down
	if err := g.SetKeybinding("side", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}

	// Side enter
	if err := g.SetKeybinding("side", gocui.KeyEnter, gocui.ModNone, onSideEnter); err != nil {
		return err
	}

	// Command enter
	if err := g.SetKeybinding("command", gocui.KeyEnter, gocui.ModNone, executeCommand); err != nil {
		return err
	}


  /*  I need to add some better functionality for this. Config files for now.

  if err := g.SetKeybinding("side", 'n', gocui.ModNone, newProfileNameView); err != nil {
    return err
  }

  if err := g.SetKeybinding("profile_name", gocui.KeyEnter, gocui.ModNone, newProfilePathView); err != nil {
    return err
  }

  if err := g.SetKeybinding("profile_path", gocui.KeyEnter, gocui.ModNone, saveNewProfile); err != nil {
    return err
  } */

	return nil
}
