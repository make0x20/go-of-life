package tgrid

import (
	"errors"
	"fmt"
	"os"
)

type Grid struct {
	Width, Height int
	Buffer        []rune
	CurrentBuffer []rune
}

func NewGrid(w, h int) Grid {
	b := make([]rune, w*h)
	cb := make([]rune, w*h)

	grid := Grid{
		Width:  w,
		Height: h,
		Buffer: b,
		CurrentBuffer: cb,
	}

	return grid
}

func (g *Grid) Clear() {
	// Clear the screen
	fmt.Fprintf(os.Stdout, "\x1b[2J\x1b[H")

	// Initialize buffer with zero values
	for i := 0; i < len(g.Buffer); i++ {
		g.Buffer[i] = 0
	}
}

func (g *Grid) FlushBuffer() {
	defer copy(g.CurrentBuffer, g.Buffer) // Create a snapshot after flushing the buffer

	g.ForEach(func(x, y, i int) {
		if g.CurrentBuffer[i] != g.Buffer[i] {
			// Move cursor to position
			fmt.Fprintf(os.Stdout, "\x1b[%d;%dH%c", y+1, x+1, g.Buffer[i])
		}
	})
}

// GetIndex gets the buffer array index from coordinates
func (g *Grid) GetIndex(x, y int) (int, error) {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return 0, errors.New("accessing non-existent grid index")
	}
	return y*g.Width + x, nil
}

// SetIndexValue sets a value at a grid array index
func (g *Grid) SetIndexValue(v rune, i int) error {
	if i < 0 || i >= len(g.Buffer) {
		return errors.New("writing non-existent grid index")
	}
	g.Buffer[i] = v

	return nil
}

func (g *Grid) GetIndexValue(i int) (rune, error) {
	if i < 0 || i >= len(g.Buffer) {
		return 0, errors.New("accessing value on non-existent grid index")
	}
	return g.Buffer[i], nil
}

// SetValue sets a value at grid coordinates
func (g *Grid) SetValue(v rune, x, y int) error {
	i, err := g.GetIndex(x, y)
	if err != nil {
		return err
	}
	g.SetIndexValue(v, i)

	return nil
}

// GetValue gets the value from grid coordinates
func (g *Grid) GetValue(x, y int) (rune, error) {
	i, err := g.GetIndex(x, y)
	if err != nil {
		return 0, err
	}

	return g.GetIndexValue(i)
}

// Grid iterator
func (g *Grid) ForEach(fn func(x, y, index int)) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			index := y*g.Width + x
			fn(x, y, index)
		}
	}
}
