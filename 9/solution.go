package main

import (
	"fmt"
	"sort"

	"github.com/unlikenesses/utils"
)

var grid [][]Point

func main() {
	lines := utils.ReadInput()
	grid = parseGrid(lines)

	part1, lowPoints := partOne()
	fmt.Println(part1)

	part2 := partTwo(lowPoints)
	fmt.Println(part2)
}

func parseGrid(lines []string) [][]Point {
	var rows [][]Point
	for y, line := range lines {
		var cols []Point
		for x, i := range line {
			point := Point{Coord{x, y}, int(i - '0')}
			cols = append(cols, point)
		}
		rows = append(rows, cols)
	}

	return rows
}

type Coord struct {
	x int
	y int
}

type Point struct {
	coord Coord
	value int
}

func partOne() (int, []Point) {
	var sum int
	var lowPoints []Point
	for _, row := range grid {
		for _, point := range row {
			if isLowPoint(point) {
				riskLevel := point.value + 1
				sum += riskLevel
				lowPoints = append(lowPoints, point)
			}
		}
	}

	return sum, lowPoints
}

func isLowPoint(point Point) bool {
	isLowPoint := true
	adjacent := getAdjacentPoints(point.coord)
	for _, a := range adjacent {
		if a.value <= point.value {
			isLowPoint = false
		}
	}

	return isLowPoint
}

func getAdjacentPoints(coord Coord) []Point {
	x := coord.x
	y := coord.y
	// Corners
	if x == 0 && y == 0 {
		return []Point{
			getPointAt(x+1, y),
			getPointAt(x, y+1),
		}
	}
	if x == 0 && y == len(grid)-1 {
		return []Point{
			getPointAt(x+1, y),
			getPointAt(x, y-1),
		}
	}
	if x == len(grid[0])-1 && y == 0 {
		return []Point{
			getPointAt(x, y+1),
			getPointAt(x-1, y),
		}
	}
	if x == len(grid[0])-1 && y == len(grid)-1 {
		return []Point{
			getPointAt(x, y-1),
			getPointAt(x-1, y),
		}
	}
	// Edges
	if x == 0 {
		return []Point{
			getPointAt(x, y-1),
			getPointAt(x, y+1),
			getPointAt(x+1, y),
		}
	}
	if x == len(grid[0])-1 {
		return []Point{
			getPointAt(x, y-1),
			getPointAt(x, y+1),
			getPointAt(x-1, y),
		}
	}
	if y == 0 {
		return []Point{
			getPointAt(x-1, y),
			getPointAt(x+1, y),
			getPointAt(x, y+1),
		}
	}
	if y == len(grid)-1 {
		return []Point{
			getPointAt(x-1, y),
			getPointAt(x+1, y),
			getPointAt(x, y-1),
		}
	}
	// Anything else
	return []Point{
		getPointAt(x-1, y),
		getPointAt(x+1, y),
		getPointAt(x, y-1),
		getPointAt(x, y+1),
	}
}

func getPointAt(x int, y int) Point {
	for gridY, row := range grid {
		for gridX, point := range row {
			if gridX == x && gridY == y {
				return point
			}
		}
	}

	return Point{}
}

func partTwo(lowPoints []Point) int {
	// Yes inefficient
	var sizes []int
	for _, point := range lowPoints {
		basin := []Point{point}
		basin = getBasin(point, basin)
		size := len(uniquePoints(basin))
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)
	biggestThree := sizes[len(sizes)-3:]

	return biggestThree[0] * biggestThree[1] * biggestThree[2]
}

func getBasin(point Point, basin []Point) []Point {
	var checked []Point
	adjacent := getAdjacentPoints(point.coord)
	filtered := filterAlreadyCheckedPoints(adjacent, checked)

	for _, a := range filtered {
		if a.value == 9 {
			continue
		}
		if a.value > point.value {
			basin = append(basin, a)
			basin = getBasin(a, basin)
		}
	}

	return basin
}

func filterAlreadyCheckedPoints(adjacent []Point, checked []Point) []Point {
	var filtered []Point
	for _, a := range adjacent {
		if !containsPoint(checked, a) {
			filtered = append(filtered, a)
		}
	}

	return filtered
}

func containsPoint(points []Point, point Point) bool {
	for _, p := range points {
		if p == point {
			return true
		}
	}

	return false
}

func uniquePoints(points []Point) []Point {
	var unique []Point
	for _, p := range points {
		if !containsPoint(unique, p) {
			unique = append(unique, p)
		}
	}

	return unique
}
