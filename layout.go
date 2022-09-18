package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

const (
  StatusView = "status"
  SidebarView = "side"
  ServersView = "servers"
  FXServerView = "main"
  FXServerCmdView = "command"
)

func layout(g *gocui.Gui) error {

	maxX, maxY := g.Size()

  if statusView, err := g.SetView(StatusView, 1, 0, maxX/6+6, 2); err != nil {
    if err != gocui.ErrUnknownView {
      return err
    }

    statusView.Title = "Status"
    statusView.FgColor = gocui.ColorRed
    fmt.Fprintln(statusView, "No server running")
  }

	if s, err := g.SetView(SidebarView, 1, 4, maxX/4, maxY-1); err != nil {
		if err != gocui.ErrUnknownView{
			return err
		}

    s.Title = "Sidebar"

    //sideX, sideY := s.Size()
    if serversView, err := g.SetView(ServersView, 1, 4, maxX/6+6, 8); err != nil {
      if err != gocui.ErrUnknownView{
        return err
      }


      serversView.FgColor = gocui.ColorWhite
      serversView.SelFgColor = gocui.ColorGreen
      serversView.Highlight = true
      serversView.Title = "Servers"
      g.Cursor = true


      for key := range Servers {
        fmt.Fprintln(serversView, key)
      }
    }

		if _, err := g.SetCurrentView(ServersView); err != nil {
			return err
		}
	}


	if v, err := g.SetView(FXServerView, int(0.2*float32(maxX)), 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "FXServer"
		v.Autoscroll = true
		v.Wrap = true

		fmt.Fprintln(v, "")
	}

	if c, err := g.SetView(FXServerCmdView, int(0.2*float32(maxX)), maxY-4, maxX-1, maxY-1); err != nil {
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
