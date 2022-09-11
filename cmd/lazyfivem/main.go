package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"sync"

	"github.com/jroimartin/gocui"
)

var mu sync.Mutex
var Scanner bufio.Scanner
var CfxCmd *exec.Cmd
var Writer io.WriteCloser

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

  if err !=  nil {
    return err
  }


  return nil
}


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

  return nil
}

func startFxServer(profile string, g *gocui.Gui) {
  if profile == "NPWD Server" {
    CfxCmd = exec.Command("X:/ServerFX/starter.bat")
    Writer, _ = CfxCmd.StdinPipe()

    r, _ := CfxCmd.StdoutPipe()

    CfxCmd.Stderr = CfxCmd.Stdout
    Scanner = *bufio.NewScanner(r)

    go listen(g, &Scanner)

    err := CfxCmd.Start()
    if err != nil {
      panic(err)
    }
  }

}

func main() {

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	defer g.Close()

  g.Cursor = true

  g.BgColor = gocui.ColorBlack
  g.FgColor = gocui.ColorWhite
  g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

  if err := keybinds(g); err != nil {
    log.Panicln(err)
  }



	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}


func layout(g *gocui.Gui) error {

	maxX, maxY := g.Size()
	if s, err := g.SetView("side", 0, 0, maxX/4, maxY - 1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

    s.Title = "Servers"
    s.FgColor = gocui.ColorWhite
    s.SelFgColor = gocui.ColorGreen
    s.Highlight = true
    
    fmt.Fprintln(s, "NPWD Server")
		fmt.Fprintln(s, "ESX Server")
		fmt.Fprintln(s, "QB Server")

    if _, err := g.SetCurrentView("side"); err != nil {
      return err
    }
	}


	if v, err := g.SetView("main", int(0.2*float32(maxX)), 0, maxX - 1, maxY - 1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

    v.Title = "FXServer"
    v.Autoscroll = true
    v.Wrap = true
    
    fmt.Fprintln(v, "")

	}

  if c, err := g.SetView("command", int(0.2*float32(maxX)), maxY - 4, maxX - 1, maxY - 1); err != nil {
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

func listen(g *gocui.Gui, scanner *bufio.Scanner) {
  for scanner.Scan() {
    line := scanner.Text()

    mu.Lock()
    updateFxServer(g, line)
  }
}

func updateFxServer(g *gocui.Gui, line string) {
  g.Update(func(g *gocui.Gui) error {
    v, err := g.View("main")
    if err != nil {
      return err
    }

    fmt.Fprintln(v, line)
    mu.Unlock()
    return nil
  })
}
