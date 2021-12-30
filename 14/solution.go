package main

import (
	"fmt"
	"regexp"

	"github.com/unlikenesses/utils"
)

type Instruction struct {
	pair    string
	element string
}

var template string
var instructions []Instruction
var pairMap = make(map[string]int)
var letterMap = make(map[string]int)

func main() {
	lines := utils.ReadInput()
	template, instructions = splitInput(lines)

	result := partOne(10)
	fmt.Println(result)

	result = partOne(40)
	fmt.Println(result)
}

func splitInput(lines []string) (string, []Instruction) {
	var template string
	var instructions []Instruction

	for i, line := range lines {
		if i == 0 {
			template = line
		} else if i > 1 {
			instructions = append(instructions, parseInstruction(line))
		}
	}

	return template, instructions
}

func parseInstruction(instruction string) Instruction {
	var parsed Instruction
	re := regexp.MustCompile(`([A-Z]{2}) -> ([A-Z])`)
	matches := re.FindStringSubmatch(instruction)
	if len(matches) > 0 {
		pair := matches[1]
		element := matches[2]
		parsed = Instruction{pair, element}
	} else {
		fmt.Println("No matches for instruction ", instruction)
	}

	return parsed
}

func partOne(steps int) int {
	pairs := splitTemplateIntoPairs(template)
	pairMap = make(map[string]int)
	letterMap = make(map[string]int)
	for _, t := range template {
		letterMap[string(t)]++
	}
	for _, p := range pairs {
		pairMap[p]++
	}

	for i := 0; i < steps; i++ {
		growPairs()
	}

	mostCommon, leastCommon := countPairMap()

	return int(mostCommon - leastCommon)
}

func countPairMap() (int, int) {
	var mostCommon, leastCommon, i int
	for _, value := range letterMap {
		if i == 0 {
			leastCommon = value
		}
		if value > mostCommon {
			mostCommon = value
		}
		if value < leastCommon {
			leastCommon = value
		}
		i++
	}

	return mostCommon, leastCommon
}

func splitTemplateIntoPairs(template string) []string {
	var pairs []string
	for i := 0; i < len(template)-1; i++ {
		pairs = append(pairs, template[i:i+2])
	}

	return pairs
}

func growPairs() {
	var tempMap = copyMap(pairMap)
	for pair, qty := range tempMap {
		if qty > 0 {
			grown := performInstruction(pair)
			newPair1 := string(grown[0]) + string(grown[1])
			newPair2 := string(grown[1]) + string(grown[2])
			letterMap[string(grown[1])] += qty
			pairMap[pair] -= qty
			pairMap[newPair1] += qty
			pairMap[newPair2] += qty
		}
	}
}

func copyMap(m map[string]int) map[string]int {
	var c = make(map[string]int)
	for k, v := range m {
		c[k] = v
	}

	return c
}

func performInstruction(pair string) string {
	instruction := findInstruction(pair)
	var result string
	result = string(pair[0]) + instruction.element + string(pair[1])

	return result
}

func findInstruction(pair string) Instruction {
	for _, instruction := range instructions {
		if instruction.pair == pair {
			return instruction
		}
	}

	return Instruction{}
}
