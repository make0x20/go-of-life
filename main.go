package main

import (
	"os"
	"go-of-life/tgrid"
	"go-of-life/gameoflife"
	"go-of-life/terminal"
)

func main() {
	gr := ParseFlags()

	terminal.SetRawMode()
	terminal.HideCursor()
	defer func() {
		terminal.ShowCursor()
		terminal.RestoreMode()
	}()

	w, h, err := terminal.GetSize()
	if err != nil {
		panic(err)
	}

	grid := tgrid.NewGrid(w, h)

	keypress := make(chan byte)

	// Listen for keypresses
	go func() {
		for {
			var input [1]byte
			os.Stdin.Read(input[:])
			keypress <- input[0]
		}
	}()

	// Initial clear
	grid.Clear()

	gameoflife.Run(&grid, gr)

	for {
		// Handle keyboard keys
		select {
		case key := <-keypress:
			if key == 'q' {
				return
			}
		default:
			// Keep doing it
		}
	}
}

