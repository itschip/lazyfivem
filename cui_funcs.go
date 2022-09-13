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

func saveNewProfile(g *gocui.Gui, v *gocui.View) error {
  var err error 
  v.Rewind()
  if err != nil {
    log.Fatal(err)
  }

  if v.Buffer() != "" {
    ServerPath = v.Buffer()

    err = g.DeleteView("profile_path")
    if err != nil {
      return err
    }
  }

  if _, err := g.SetCurrentView("side"); err != nil {
    return err
  }
  return nil
}

func newProfileNameView(g *gocui.Gui, v *gocui.View) error {
  maxX, maxY := g.Size()

  if vp_n, err := g.SetView("profile_name", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2);  err != nil {
    if err != gocui.ErrUnknownView {
      return err
    }

    vp_n.Title = "Add server name - Enter to continue | Abort with q "
    vp_n.BgColor = gocui.ColorBlack
    vp_n.Editable = true
  }


  if _, err := g.SetCurrentView("profile_name"); err != nil {
    return err
  }


  return nil
}

func newProfilePathView(g *gocui.Gui, v *gocui.View) error {
  var err error 
  v.Rewind()
  if err != nil {
    log.Fatal(err)
  }

	if v.Buffer() != "" {
    ServerName = v.Buffer()
  }

  maxX, maxY := g.Size()

  if vp_p, err := g.SetView("profile_path", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2);  err != nil {
    if err != gocui.ErrUnknownView {
      return err
    }

    vp_p.Title = "Add startup path - Enter to finish | Abort with q "
    vp_p.BgColor = gocui.ColorBlack
    vp_p.Editable = true
  }

  if _, err := g.SetCurrentView("profile_path"); err != nil {
    return err
  }

  // delete profile_name view
  err = g.DeleteView("profile_name")
  if err != nil {
    return err
  }
  
  return nil
}
