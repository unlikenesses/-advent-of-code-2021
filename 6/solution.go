package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/unlikenesses/utils"
)

func main() {
	lines := utils.ReadInput()
	fish := parseFish(lines[0])

	part1 := partOneAndTwo(fish, 80)
	fmt.Println(part1)

	fish = parseFish(lines[0])

	part2 := partOneAndTwo(fish, 256)
	fmt.Println(part2)
}

func parseFish(lines string) map[int]int {
	m := make(map[int]int)
	timers := strings.Split(lines, ",")
	for _, timer := range timers {
		timeInt, _ := strconv.Atoi(timer)
		m[timeInt] += 1
	}

	return m
}

func partOneAndTwo(fish map[int]int, days int) int {
	var toSpawn int
	for d := 0; d < days; d++ {
		for t := 0; t < 9; t++ {
			if t == 0 {
				toSpawn = fish[t]
			}
			if t < 8 {
				fish[t] = fish[t+1]
			}
			if t == 8 {
				fish[t] = toSpawn
				fish[6] += toSpawn
			}
		}
	}
	var sum int
	for _, numFish := range fish {
		sum += numFish
	}

	return sum
}
