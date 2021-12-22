package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/unlikenesses/utils"
)

var incomplete []string

var matches = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
	'<': '>',
}

func main() {
	lines := utils.ReadInput()

	part1 := partOne(lines)
	fmt.Println(part1)

	part2 := partTwo()
	fmt.Println(part2)
}

func partOne(lines []string) int {
	var sum int
	var corrupted []string

	for _, line := range lines {
		lineCorrupted, score := isCorrupted(line)
		if lineCorrupted {
			corrupted = append(corrupted, line)
			sum += score
		} else {
			incomplete = append(incomplete, line)
		}
	}

	return sum
}

func isCorrupted(line string) (bool, int) {
	var openers []rune
	for _, char := range line {
		if isOpener(char) {
			openers = append(openers, char)
		} else if isCloser(char) {
			numOpeners := len(openers)
			lastOpener := openers[numOpeners-1]
			if doMatch(lastOpener, char) {
				// Remove this opener from the list:
				openers = openers[:numOpeners-1]
			} else {
				return true, getCorruptedCharScore(char)
			}
		} else {
			fmt.Printf("Char in line %s not categorisable", line)
			os.Exit(1)
		}
	}

	return false, 0
}

func isOpener(char rune) bool {
	openers := []rune{
		'(',
		'{',
		'[',
		'<',
	}
	return utils.InCharArray(openers, char)
}

func isCloser(char rune) bool {
	closers := []rune{
		')',
		'}',
		']',
		'>',
	}
	return utils.InCharArray(closers, char)
}

func doMatch(opener rune, closer rune) bool {
	return matches[opener] == closer
}

func getCorruptedCharScore(char rune) int {
	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	return scores[char]
}

func partTwo() int {
	var scores []int
	for _, line := range incomplete {
		score := getCompletionStringScore(line)
		scores = append(scores, score)
	}
	sort.Ints(scores)

	return scores[len(scores)/2]
}

func getCompletionStringScore(line string) int {
	var score int
	var openers []rune
	for _, char := range line {
		if isOpener(char) {
			openers = append(openers, char)
		} else if isCloser(char) {
			numOpeners := len(openers)
			lastOpener := openers[numOpeners-1]
			if doMatch(lastOpener, char) {
				// Remove this opener from the list:
				openers = openers[:numOpeners-1]
			} else {
				fmt.Printf("Error: Found corrupted line %s", line)
				os.Exit(1)
			}
		} else {
			fmt.Printf("Char in line %s not categorisable", line)
			os.Exit(1)
		}
	}

	for o := len(openers) - 1; o >= 0; o-- {
		opener := openers[o]
		closer := matches[opener]
		score = score * 5
		score += getCompletingCharScore(closer)
	}

	return score
}

func getCompletingCharScore(char rune) int {
	scores := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	return scores[char]
}
