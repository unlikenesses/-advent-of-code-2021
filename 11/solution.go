package main

import (
	"fmt"

	"github.com/unlikenesses/utils"
)

var grid [][]Octopus

type Coord struct {
	x int
	y int
}

type Octopus struct {
	energy  int
	flashed bool
}

func main() {
	lines := utils.ReadInput()
	grid = parseGrid(lines)

	part1 := partOne()
	fmt.Println(part1)

	grid = parseGrid(lines)
	part2 := partTwo()
	fmt.Println(part2)
}

func parseGrid(lines []string) [][]Octopus {
	var grid [][]Octopus
	for _, line := range lines {
		var row []Octopus
		for _, char := range line {
			val := int(char - '0')
			octopus := Octopus{val, false}
			row = append(row, octopus)
		}
		grid = append(grid, row)
	}

	return grid
}

func partOne() int {
	var totalFlashes int

	for s := 0; s < 100; s++ {
		numFlashes := mutateGrid()
		totalFlashes += numFlashes
	}

	return totalFlashes
}

func partTwo() int {
	var numFlashes int
	var step int

	gridSize := len(grid) * len(grid[0])
	for numFlashes < gridSize {
		numFlashes = mutateGrid()
		step++
	}

	return step
}

func mutateGrid() int {
	// Increase all by one:
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			grid[y][x].energy += 1
			grid[y][x].flashed = false
		}
	}
	// Check for > 9:
	flashes := checkFlashes()
	// Set flashed to zero:
	for _, flashed := range flashes {
		grid[flashed.y][flashed.x].energy = 0
	}

	return len(flashes)
}

func checkFlashes() []Coord {
	foundImminentFlash := true
	var flashes []Coord
	for foundImminentFlash {
		foundImminentFlash = false
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				val := grid[y][x].energy
				if val > 9 && grid[y][x].flashed == false {
					grid[y][x].flashed = true
					flash(x, y)
					flashes = append(flashes, Coord{x, y})
					foundImminentFlash = true
				}
			}
		}
	}

	return flashes
}

func flash(x, y int) {
	// top row:
	if y > 0 {
		grid[y-1][x].energy += 1
		if x > 0 {
			grid[y-1][x-1].energy += 1
		}
		if x < len(grid[0])-1 {
			grid[y-1][x+1].energy += 1
		}
	}
	// middle row:
	if x > 0 {
		grid[y][x-1].energy += 1
	}
	if x < len(grid[0])-1 {
		grid[y][x+1].energy += 1
	}
	// bottom row:
	if y < len(grid)-1 {
		grid[y+1][x].energy += 1
		if x > 0 {
			grid[y+1][x-1].energy += 1
		}
		if x < len(grid[0])-1 {
			grid[y+1][x+1].energy += 1
		}
	}
}
