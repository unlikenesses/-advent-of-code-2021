package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/unlikenesses/utils"
)

type Entry struct {
	signal []string
	output []string
}

func main() {
	lines := utils.ReadInput()
	entries := parseEntries(lines)

	part1 := partOne(entries)
	fmt.Println(part1)

	part2 := partTwo(entries)
	fmt.Println(part2)
}

func parseEntries(lines []string) []Entry {
	var entries []Entry
	re := regexp.MustCompile(`([abcdefg]{1,7}) ([abcdefg]{1,7}) ([abcdefg]{1,7}) ([abcdefg]{1,7}) ([abcdefg]{1,7}) ([abcdefg]{1,7}) ([abcdefg]{1,7}) ([abcdefg]{1,7}) ([abcdefg]{1,7}) ([abcdefg]{1,7}) \| ([abcdefg]{1,7}) ([abcdefg]{1,7}) ([abcdefg]{1,7}) ([abcdefg]{1,7})`)
	for _, line := range lines {
		var signal []string
		var output []string
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			for s := 1; s <= 10; s++ {
				signal = append(signal, matches[s])
			}
			for s := 11; s <= 14; s++ {
				output = append(output, matches[s])
			}
			entries = append(entries, Entry{signal, output})
		} else {
			fmt.Println("No matches for line ", line)
		}
	}

	return entries
}

func partOne(entries []Entry) int {
	var count int
	for _, entry := range entries {
		for _, o := range entry.output {
			l := len(o)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count++
			}
		}
	}

	return count
}

func partTwo(entries []Entry) int {
	sum := 0
	for _, entry := range entries {
		sum += calculateDisplay(entry)
	}

	return sum
}

func calculateDisplay(entry Entry) int {
	mapping := calculateMapping(entry.signal)
	originalMapping := map[string]string{
		"abcefg":  "0",
		"cf":      "1",
		"acdeg":   "2",
		"acdfg":   "3",
		"bcdf":    "4",
		"abdfg":   "5",
		"abdefg":  "6",
		"acf":     "7",
		"abcdefg": "8",
		"abcdfg":  "9",
	}
	var digits string
	for _, o := range entry.output {
		translated := translateOutput(o, mapping)
		sorted := sortOutput(translated)
		digits += originalMapping[sorted]
	}
	result, _ := strconv.Atoi(digits)

	return result
}

func calculateMapping(signals []string) map[string]string {
	m := make(map[string]string)
	// Calculate c & f:
	s2 := getSignalsOfLength(signals, 2)
	temp1 := string(s2[0][0])
	temp2 := string(s2[0][1])
	// Only one of the 6-size signals doesn't have both c & f in it
	s6 := getSignalsOfLength(signals, 6)
	// Jesus there has to be a better way of doing it than this
	for _, s := range s6 {
		if strings.Contains(s, temp1) && !strings.Contains(s, temp2) {
			m[temp1] = "f"
			m[temp2] = "c"
		} else if !strings.Contains(s, temp1) && strings.Contains(s, temp2) {
			m[temp1] = "c"
			m[temp2] = "f"
		}
	}
	// Calculate a
	s3 := getSignalsOfLength(signals, 3)
	temp3 := ""
	for _, s := range s3[0] {
		_, ok := m[string(s)]
		if !ok {
			m[string(s)] = "a"
			temp3 = string(s)
		}
	}
	// Calculate b & d
	s4 := getSignalsOfLength(signals, 4)
	bAndD := strings.Replace(s4[0], temp1, "", 1)
	bAndD = strings.Replace(bAndD, temp2, "", 1)
	// Only one of the 6-size signals doesn't have both b & d in it
	temp4 := string(bAndD[0])
	temp5 := string(bAndD[1])
	for _, s := range s6 {
		if strings.Contains(s, temp4) && !strings.Contains(s, temp5) {
			m[temp4] = "b"
			m[temp5] = "d"
		} else if !strings.Contains(s, temp4) && strings.Contains(s, temp5) {
			m[temp4] = "d"
			m[temp5] = "b"
		}
	}
	// Calculate e & g
	s7 := getSignalsOfLength(signals, 7)
	eAndG := strings.Replace(s7[0], temp1, "", 1)
	eAndG = strings.Replace(eAndG, temp2, "", 1)
	eAndG = strings.Replace(eAndG, temp3, "", 1)
	eAndG = strings.Replace(eAndG, temp4, "", 1)
	eAndG = strings.Replace(eAndG, temp5, "", 1)
	// Only one of the 6-size signals doesn't have both b & g in it
	temp6 := string(eAndG[0])
	temp7 := string(eAndG[1])
	for _, s := range s6 {
		if strings.Contains(s, temp6) && !strings.Contains(s, temp7) {
			m[temp6] = "g"
			m[temp7] = "e"
		} else if !strings.Contains(s, temp6) && strings.Contains(s, temp7) {
			m[temp6] = "e"
			m[temp7] = "g"
		}
	}

	return m
}

func getSignalsOfLength(signals []string, length int) []string {
	var matching []string
	for _, signal := range signals {
		if len(signal) == length {
			matching = append(matching, signal)
		}
	}

	return matching
}

func translateOutput(output string, mapping map[string]string) string {
	var translated string
	for _, o := range output {
		translated += mapping[string(o)]
	}

	return translated
}

func sortOutput(output string) string {
	s := strings.Split(output, "")
	sort.Strings(s)

	return strings.Join(s, "")
}
