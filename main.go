package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"go-of-life/tgrid"
	"go-of-life/gameoflife"
	"go-of-life/terminal"
)

func main() {
	gr := ParseFlags()

	// Start pprof server for profiling
	go func() {
		log.Println("Starting pprof server on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

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

