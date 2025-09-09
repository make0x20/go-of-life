package gameoflife

import (
	"math/rand"
	"slices"
	"go-of-life/tgrid"
	"time"
)

type Rules struct {
	NeighborhoodSize int
	Speed            time.Duration
	Cell             rune
	PopulateChance   float64
	SurvivalCounts   []int
	BirthCounts      []int
}

func Run(grid *tgrid.Grid, r Rules) chan int {
	keyChannel := make(chan int)

	// Populate withrandom cells
	populate(grid, r.Cell, r.PopulateChance)

	grid.FlushBuffer()

	// Pause for a second just to give the user a quick view of the initial state
	time.Sleep(time.Second)

	go func() {
		for {
			select {
			case key := <-keyChannel:
				if key == 0 {
					close(keyChannel)
					return
				}

			default:
				iterate(grid, r)
			}

			grid.FlushBuffer()

			time.Sleep(r.Speed)
		}
	}()

	return keyChannel
}

func iterate(grid *tgrid.Grid, r Rules) {
	// Create a copy of the buffer for new iteration
	newBuffer := make([]rune, len(grid.Buffer))
	copy(newBuffer, grid.Buffer)

	grid.ForEachAsync(func(x, y, i int) {
		currentCell, _ := grid.GetValue(x, y)
		neighbors := checkNeighborhood(x, y, grid.Width, grid.Height, r.Cell, grid, r)

		if currentCell == r.Cell { // Currently alive
			if slices.Contains(r.SurvivalCounts, neighbors) {
				// stays alive
			} else {
				newBuffer[i] = ' ' // dies
			}
		} else { // Currently dead
			if slices.Contains(r.BirthCounts, neighbors) {
				newBuffer[i] = r.Cell // becomes alive
			}
		}
	})

	// Override the buffer with new iteration
	grid.Buffer = newBuffer
}

func checkNeighborhood(x, y, width, height int, v rune, grid *tgrid.Grid, r Rules) int {
	alive := 0
	nSize := r.NeighborhoodSize

	for nx := x - nSize; nx <= x+nSize; nx++ {
		for ny := y - nSize; ny <= y+nSize; ny++ {
			// Skip the center cell - x, y
			if nx == x && ny == y {
				continue
			}

			// Skip if outside of grid
			if isOutOfBounds(nx, ny, width, height) {
				continue
			}

			// If cell is alive add to alive count
			index := ny*width + nx
			nv := grid.Buffer[index]

			if nv == v {
				alive++
			}
		}
	}

	return alive
}

// Check if a cell is outside of the grid
func isOutOfBounds(x, y, w, h int) bool {
	return x < 0 || x >= w || y < 0 || y >= h
}


// Populate the grid with random cells
func populate(grid *tgrid.Grid, cell rune, populateChance float64) {
	grid.ForEach(func(x, y, index int) {
		if rand.Float64() < populateChance {
			grid.SetValue(cell, x, y)
		} else {
			grid.SetValue(' ', x, y)
		}
	})
}
