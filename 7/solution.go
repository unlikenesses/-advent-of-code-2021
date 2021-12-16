package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/unlikenesses/utils"
)

func main() {
	lines := utils.ReadInput()
	positions, max := parsePositions(lines[0])

	part1 := partOne(positions, max)
	fmt.Println(part1)

	part2 := partTwo(positions, max)
	fmt.Println(part2)
}

func parsePositions(lines string) ([]int, int) {
	posArray := strings.Split(lines, ",")
	var positions []int
	var max int
	for _, pos := range posArray {
		posInt, _ := strconv.Atoi(pos)
		if posInt > max {
			max = posInt
		}
		positions = append(positions, posInt)
	}

	return positions, max
}

func partOne(positions []int, max int) int {
	var minFuel int
	for i := 0; i < max; i++ {
		fuel := 0
		for _, p := range positions {
			diff := utils.AbsInt(p - i)
			fuel += diff
		}
		if i == 0 {
			minFuel = fuel
		} else if minFuel > fuel {
			minFuel = fuel
		}
	}

	return minFuel
}

func partTwo(positions []int, max int) int {
	var minFuel int
	for i := 0; i < max; i++ {
		fuel := 0
		for _, p := range positions {
			diff := utils.AbsInt(p - i)
			fuel += sumAll(diff)
		}
		if i == 0 {
			minFuel = fuel
		} else if minFuel > fuel {
			minFuel = fuel
		}
	}

	return minFuel
}

func sumAll(num int) int {
	sum := 0
	for n := 1; n <= num; n++ {
		sum += n
	}

	return sum
}
