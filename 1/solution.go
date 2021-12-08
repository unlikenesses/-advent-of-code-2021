package main

import (
	"fmt"
	"strconv"

	"github.com/unlikenesses/utils"
)

func main() {
	lines := utils.ReadInput()

	part1 := partOne(lines)
	fmt.Println(part1)

	part2 := partTwo(lines)
	fmt.Println(part2)
}

func partOne(lines []string) int {
	prev := 0
	total := 0
	for i, line := range lines {
		depth, _ := strconv.Atoi(line)
		if i == 0 {
			prev = depth
			continue
		}
		if depth > prev {
			total++
		}
		prev = depth
	}
	return total
}

func partTwo(lines []string) int {
	window_sums := getWindowSums(lines)
	prev := 0
	total := 0
	for i, sum := range window_sums {
		if i == 0 {
			prev = sum
			continue
		}
		if sum > prev {
			total++
		}
		prev = sum
	}

	return total
}

func getWindowSums(lines []string) []int {
	size := len(lines)
	var windows []int
	for i := 0; i < size-2; i++ {
		depth1, _ := strconv.Atoi(lines[i])
		depth2, _ := strconv.Atoi(lines[i+1])
		depth3, _ := strconv.Atoi(lines[i+2])
		window := []int{depth1, depth2, depth3}
		windowSum := utils.SumIntArray(window)
		windows = append(windows, windowSum)
	}

	return windows
}
