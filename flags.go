package main

import (
	"flag"
	"fmt"
	"go-of-life/gameoflife"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseIntSlice(s string) []int {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	result := make([]int, len(parts))
	for i, part := range parts {
		val, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid number in list: %s\n", part)
			os.Exit(1)
		}
		result[i] = val
	}
	return result
}

func getDefaultRules() gameoflife.Rules {
	return gameoflife.Rules{
		NeighborhoodSize: 1,
		Cell:            '█',
		Speed:           50 * time.Millisecond,
		PopulateChance:  0.2,
		SurvivalCounts:  []int{2, 3},
		BirthCounts:     []int{3},
	}
}

func ParseFlags() gameoflife.Rules {
	var neighborhoodSize = flag.Int("neighborhood", 0, "Neighborhood size (default: 1 for Conway's Game of Life)")
	var cell = flag.String("cell", "", "Cell character (default: █)")
	var speed = flag.Int("speed", 0, "Speed in milliseconds (default: 50)")
	var populateChance = flag.Float64("populate", 0.0, "Population chance 0.0-1.0 (default: 0.2)")
	var survivalCounts = flag.String("survival", "", "Comma-separated survival neighbor counts (default: 2,3)")
	var birthCounts = flag.String("birth", "", "Comma-separated birth neighbor counts (default: 3)")
	
	flag.Parse()

	gr := getDefaultRules()
	
	if *neighborhoodSize > 0 {
		gr.NeighborhoodSize = *neighborhoodSize
	}
	
	if *cell != "" {
		cellRunes := []rune(*cell)
		if len(cellRunes) > 0 {
			gr.Cell = cellRunes[0]
		}
	}
	
	if *speed > 0 {
		gr.Speed = time.Duration(*speed) * time.Millisecond
	}
	
	if *populateChance > 0.0 {
		gr.PopulateChance = *populateChance
	}
	
	if *survivalCounts != "" {
		gr.SurvivalCounts = parseIntSlice(*survivalCounts)
	}
	
	if *birthCounts != "" {
		gr.BirthCounts = parseIntSlice(*birthCounts)
	}

	return gr
}
