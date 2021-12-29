package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/unlikenesses/utils"
)

var grid [][]int

func main() {
	lines := utils.ReadInput()
	coords, instructions := splitInput(lines)
	parseGrid(coords, instructions)

	partOne(instructions)
}

func splitInput(lines []string) ([]string, []string) {
	var coords []string
	var instructions []string
	inCoords := true
	for _, line := range lines {
		if line == "" {
			inCoords = false
			continue
		}
		if inCoords {
			coords = append(coords, line)
		} else {
			instructions = append(instructions, line)
		}
	}

	return coords, instructions
}

func parseGrid(coords, instructions []string) {
	maxX, maxY := getGridBounds(instructions)
	// Set all grid to 0:
	grid = makeGrid(maxX, maxY)
	for _, line := range coords {
		s := strings.Split(line, ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		grid[y][x] = 1
	}
}

func getGridBounds(instructions []string) (int, int) {
	var maxX, maxY int
	for _, instruction := range instructions {
		axis, pos := parseInstruction(instruction)
		if axis == "x" {
			maxX = (pos * 2) + 1
		} else if axis == "y" {
			maxY = (pos * 2) + 1
		}
		if maxX > 0 && maxY > 0 {
			break
		}
	}

	return maxX, maxY
}

func makeGrid(width, height int) [][]int {
	var tempGrid [][]int
	for y := 0; y < height; y++ {
		var row []int
		for x := 0; x < width; x++ {
			row = append(row, 0)
		}
		tempGrid = append(tempGrid, row)
	}

	return tempGrid
}

func partOne(instructions []string) {
	for i, instruction := range instructions {
		axis, pos := parseInstruction(instruction)
		performFold(axis, pos)
		if i == 0 {
			fmt.Printf("After 1 fold, there are %d dots\n", numDots())
		}
	}

	var char string
	for _, row := range grid {
		for _, col := range row {
			char = " "
			if col == 1 {
				char = "*"
			}
			fmt.Printf(char)
		}
		fmt.Println()
	}
}

func parseInstruction(instruction string) (string, int) {
	var axis string
	var pos int
	re := regexp.MustCompile(`fold along (x|y)=(\d+)`)
	matches := re.FindStringSubmatch(instruction)
	if len(matches) > 0 {
		axis = matches[1]
		pos, _ = strconv.Atoi(matches[2])
	} else {
		fmt.Println("No matches for instruction ", instruction)
	}

	return axis, pos
}

func performFold(axis string, pos int) {
	var tempGrid, folded [][]int
	switch axis {
	case "y":
		tempGrid = foldGrid(grid, pos)
	case "x":
		tempGrid = rotateRight(grid)
		folded = foldGrid(tempGrid, pos)
		tempGrid = rotateLeft(folded)
	}

	grid = tempGrid
}

func foldGrid(g [][]int, pos int) [][]int {
	var tempGrid [][]int
	originalHeight := len(g)
	belowFoldHeight := originalHeight - pos - 1
	midPoint := float64(originalHeight / 2)

	if pos >= int(math.Floor(midPoint)) {
		newHeight := originalHeight - belowFoldHeight - 1
		overlayStart := newHeight - belowFoldHeight
		tempY := originalHeight - 1
		for h, row := range g {
			if h >= newHeight {
				break
			}

			if h >= overlayStart {
				mergedRow := mergeRows(row, g[tempY])
				tempGrid = append(tempGrid, mergedRow)
			} else {
				tempGrid = append(tempGrid, row)
			}
			tempY--
		}
	}

	return tempGrid
}

func mergeRows(r1, r2 []int) []int {
	// fmt.Println("Merging:")
	// fmt.Println(r1, r2)
	var merged []int
	var mc int
	for i, c := range r1 {
		if c == 1 || r2[i] == 1 {
			mc = 1
		} else {
			mc = 0
		}
		merged = append(merged, mc)
	}
	// fmt.Println("Merged:")
	// fmt.Println(merged)

	return merged
}

func rotateRight(g [][]int) [][]int {
	w := len(g)
	h := len(g[0])
	// fmt.Println("Rotating", g)
	rotated := make([][]int, h)
	for x := 0; x < h; x++ {
		rotated[x] = make([]int, w)
	}
	// fmt.Println("Rotated", rotated)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			rotated[i][j] = g[w-j-1][i]
		}
	}

	return rotated
}

func rotateLeft(g [][]int) [][]int {
	w := len(g)
	h := len(g[0])
	rotated := make([][]int, h)
	for x := 0; x < h; x++ {
		rotated[x] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			rotated[i][j] = g[j][h-i-1]
		}
	}

	return rotated
}

func numDots() int {
	var sum int
	for _, row := range grid {
		for _, col := range row {
			if col == 1 {
				sum++
			}
		}
	}

	return sum
}
