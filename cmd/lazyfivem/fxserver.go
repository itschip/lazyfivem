package main

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/jroimartin/gocui"
)

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
