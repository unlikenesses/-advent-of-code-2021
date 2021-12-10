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

func partOne(lines []string) int64 {
	gamma := calculateGamma(lines)
	epsilon := calculateEpsilon(gamma)

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)

	return g * e
}

func calculateGamma(lines []string) string {
	numberSize := len(lines[0])
	var gamma string
	for i := 0; i < numberSize; i++ {
		oneCount := 0
		zeroCount := 0
		for _, line := range lines {
			if line[i] == '1' {
				oneCount++
			} else {
				zeroCount++
			}
		}
		if oneCount > zeroCount {
			gamma += "1"
		} else {
			gamma += "0"
		}
	}

	return gamma
}

func calculateEpsilon(gamma string) string {
	// Dumb, I know
	var epsilon string
	for i := 0; i < len(gamma); i++ {
		if gamma[i] == '1' {
			epsilon += "0"
		} else {
			epsilon += "1"
		}
	}

	return epsilon
}

func partTwo(lines []string) int64 {
	numberSize := len(lines[0])

	ogr := calculateOGR(lines, numberSize)
	csr := calculateCSR(lines, numberSize)

	o, _ := strconv.ParseInt(ogr, 2, 64)
	c, _ := strconv.ParseInt(csr, 2, 64)

	return o * c
}

func calculateOGR(lines []string, numberSize int) string {
	for i := 0; i < numberSize; i++ {
		mostCommonBit := mostCommonBit(lines, i)
		lines = getValidLines(lines, i, mostCommonBit)
		if len(lines) < 2 {
			break
		}
	}
	return lines[0]
}

func calculateCSR(lines []string, numberSize int) string {
	for i := 0; i < numberSize; i++ {
		leastCommonBit := leastCommonBit(lines, i)
		lines = getValidLines(lines, i, leastCommonBit)
		if len(lines) < 2 {
			break
		}
	}
	return lines[0]
}

func mostCommonBit(lines []string, i int) rune {
	oneCount, zeroCount := getCount(lines, i)
	if oneCount >= zeroCount {
		return '1'
	} else {
		return '0'
	}
}

func leastCommonBit(lines []string, i int) rune {
	oneCount, zeroCount := getCount(lines, i)
	if oneCount < zeroCount {
		return '1'
	} else {
		return '0'
	}
}

func getCount(lines []string, i int) (int, int) {
	oneCount := 0
	zeroCount := 0
	for _, line := range lines {
		if line[i] == '1' {
			oneCount++
		} else {
			zeroCount++
		}
	}

	return oneCount, zeroCount
}

func getValidLines(lines []string, i int, bit rune) []string {
	var newLines []string
	for _, line := range lines {
		if string(line[i]) == string(bit) {
			newLines = append(newLines, line)
		}
	}

	return newLines
}
