package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/unlikenesses/utils"
)

type Coord struct {
	x int
	y int
}

type Vent struct {
	start Coord
	end   Coord
}

func main() {
	lines := utils.ReadInput()
	vents := parseVents(lines)

	part1 := partOne(vents, false)
	fmt.Println(part1)

	part2 := partOne(vents, true)
	fmt.Println(part2)
}

func parseVents(lines []string) []Vent {
	var vents []Vent
	re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			x1, _ := strconv.Atoi(matches[1])
			y1, _ := strconv.Atoi(matches[2])
			x2, _ := strconv.Atoi(matches[3])
			y2, _ := strconv.Atoi(matches[4])
			coords1 := Coord{x1, y1}
			coords2 := Coord{x2, y2}
			vents = append(vents, Vent{coords1, coords2})
		} else {
			fmt.Println("No matches for line ", line)
		}
	}

	return vents
}

func partOne(vents []Vent, includeDiagonals bool) int {
	coords := getAllCoordsInVents(vents, includeDiagonals)
	sum := 0
	for _, incidence := range coords {
		if incidence >= 2 {
			sum++
		}
	}

	return sum
}

func getAllCoordsInVents(vents []Vent, includeDiagonals bool) map[Coord]int {
	m := make(map[Coord]int)
	for _, vent := range vents {
		ventCoords := getAllCoordsInVent(vent, includeDiagonals)
		if len(ventCoords) > 0 {
			for _, vc := range ventCoords {
				m[vc] = m[vc] + 1
			}
		}
	}

	return m
}

func getAllCoordsInVent(vent Vent, includeDiagonals bool) []Coord {
	var coords []Coord
	if vent.start.x == vent.end.x {
		for y := utils.MinInt(vent.start.y, vent.end.y); y <= utils.MaxInt(vent.start.y, vent.end.y); y++ {
			coords = append(coords, Coord{vent.start.x, y})
		}
		return coords
	}
	if vent.start.y == vent.end.y {
		for x := utils.MinInt(vent.start.x, vent.end.x); x <= utils.MaxInt(vent.start.x, vent.end.x); x++ {
			coords = append(coords, Coord{x, vent.start.y})
		}
		return coords
	}
	if includeDiagonals {
		dx := 1
		if vent.start.x > vent.end.x {
			dx = -1
		}
		dy := 1
		if vent.start.y > vent.end.y {
			dy = -1
		}
		len := utils.AbsInt(vent.start.x-vent.end.x) + 1
		for i := 0; i < len; i++ {
			coord := Coord{vent.start.x + (i * dx), vent.start.y + (i * dy)}
			coords = append(coords, coord)
		}
	}
	return coords
}
