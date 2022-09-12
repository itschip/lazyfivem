package main

import (
	"bufio"
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
