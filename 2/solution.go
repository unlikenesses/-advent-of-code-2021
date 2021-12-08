package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/unlikenesses/utils"
)

func main() {
	lines := utils.ReadInput()

	instructions := parseInstructions(lines)

	// for _, inst := range instructions {
	// 	fmt.Println(inst)
	// }

	part1 := partOne(instructions)
	fmt.Println(part1)

	part2 := partTwo(instructions)
	fmt.Println(part2)
}

type Instruction struct {
	direction string
	magnitude int
}

func partOne(instructions []Instruction) int {
	horizontal := 0
	depth := 0
	for _, instruction := range instructions {
		magnitude := instruction.magnitude
		switch instruction.direction {
		case "forward":
			horizontal += magnitude
		case "up":
			depth -= magnitude
		case "down":
			depth += magnitude
		}
	}

	return horizontal * depth
}

func partTwo(instructions []Instruction) int {
	horizontal := 0
	depth := 0
	aim := 0
	for _, instruction := range instructions {
		magnitude := instruction.magnitude
		switch instruction.direction {
		case "forward":
			horizontal += magnitude
			depth += aim * magnitude
		case "up":
			aim -= magnitude
		case "down":
			aim += magnitude
		}
	}

	return horizontal * depth
}

func parseInstructions(lines []string) []Instruction {
	var instructions []Instruction
	re := regexp.MustCompile(`(forward|up|down) (\d)`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			direction := matches[1]
			magnitude, _ := strconv.Atoi(matches[2])
			instruction := Instruction{direction, magnitude}
			instructions = append(instructions, instruction)
		} else {
			fmt.Println("No matches for line ", line)
		}
	}

	return instructions
}
