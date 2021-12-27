package main

import (
	"fmt"
	"regexp"

	"github.com/unlikenesses/utils"
)

var connections = make(map[string][]string)

func main() {
	lines := utils.ReadInput()
	makeConnectionMap(lines)

	part1 := partOne()
	fmt.Println(part1)

	part2 := partTwo()
	fmt.Println(part2)
}

func makeConnectionMap(lines []string) {
	re := regexp.MustCompile(`(\w+)-(\w+)`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			start := matches[1]
			end := matches[2]
			if start != "end" {
				connections[start] = append(connections[start], end)
			}
			if start != "start" && end != "end" {
				connections[end] = append(connections[end], start)
			}
		} else {
			fmt.Println("No matches for line ", line)
		}
	}
}

func partOne() int {
	paths := findPath("start", 1)

	return len(paths)
}

func partTwo() int {
	paths := findPath("start", 2)

	return len(paths)
}

func findPath(start string, part int) [][]string {
	var path = []string{start}
	var paths = [][]string{path}
	var solutions [][]string

	for len(paths) > 0 {
		// Pop first path:
		path = paths[0]
		paths = paths[1:]
		lastCaveInPath := path[len(path)-1]
		if lastCaveInPath == "end" {
			solutions = append(solutions, path)
		}
		// Check connections to lastCaveInPath
		for _, conn := range connections[lastCaveInPath] {
			if (part == 1 && !visitedCavePartOne(conn, path)) ||
				(part == 2 && !visitedCavePartTwo(conn, path)) {
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, conn)
				paths = append(paths, newPath)
			}
		}
	}

	return solutions
}

func visitedCavePartOne(cave string, path []string) bool {
	if utils.IsUpperCase(cave) {
		return false
	}
	for _, c := range path {
		if c == cave {
			return true
		}
	}
	return false
}

func visitedCavePartTwo(cave string, path []string) bool {
	// We don't want to return to start
	if cave == "start" {
		return true
	}
	// We can always return to a big cave
	if utils.IsUpperCase(cave) {
		return false
	}
	aCaveHasBeenVisitedTwice := false
	caveIsInPath := false
	// Make map of visits for each small cave
	visits := make(map[string]int)
	for _, c := range path {
		if !utils.IsUpperCase(c) {
			visits[c] += 1
		}
		if c == cave {
			caveIsInPath = true
		}
	}
	// The cave hasn't been visited before:
	if !caveIsInPath {
		return false
	}
	// Has any cave been visited more than once?
	for _, numVisits := range visits {
		if numVisits > 1 {
			aCaveHasBeenVisitedTwice = true
		}
	}
	// The cave has been visited before and a cave has already had more than one visit:
	if aCaveHasBeenVisitedTwice {
		return true
	}

	// The cave has been visited before and no cave has had more than one visit:
	return false
}
