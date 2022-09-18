package main

import (
	"bufio"
	"io"
	"log"
	"os/exec"
	"sync"

	"github.com/jroimartin/gocui"
)

var (
	mu      sync.Mutex
	Scanner bufio.Scanner
	CfxCmd  *exec.Cmd
	Writer  io.WriteCloser

	ServerName = ""
	ServerPath = ""

	Servers = make(map[string]string)
)

func main() {
	getConfigValues()
	getAllServers()

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
