package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/unlikenesses/utils"
)

func main() {
	lines := utils.ReadInput()

	numbers := convertLineToIntArray(lines[0], ",")
	boards := parseBoards(lines)

	part1 := partOne(numbers, boards)
	fmt.Println(part1)

	boards = parseBoards(lines)

	part2 := partTwo(numbers, boards)
	fmt.Println(part2)
}

type Square struct {
	value   int
	crossed bool
}

type Board struct {
	rows [][]Square
}

func partOne(numbers []int, boards []Board) int {
	winner, calledNumber := getWinningBoard(boards, numbers)

	return calculateBoardScore(winner, calledNumber)
}

func partTwo(numbers []int, boards []Board) int {
	winner, calledNumber := getLastWinningBoard(boards, numbers)

	return calculateBoardScore(winner, calledNumber)
}

func parseBoards(lines []string) []Board {
	var boards []Board
	var rows [][]Square
	var squares []Square
	for i, line := range lines {
		if i > 1 {
			if len(line) > 0 {
				row := convertLineToIntArray(line, " ")
				for _, num := range row {
					squares = append(squares, Square{num, false})
				}
				rows = append(rows, squares)
				squares = nil
			} else {
				boards = append(boards, Board{rows})
				rows = nil
			}
		}
	}
	boards = append(boards, Board{rows})

	return boards
}

func convertLineToIntArray(line string, sep string) []int {
	var row []int
	split := strings.Split(line, sep)
	for _, char := range split {
		if len(char) > 0 {
			v, _ := strconv.Atoi(char)
			row = append(row, v)
		}
	}

	return row
}

func getWinningBoard(boards []Board, numbers []int) (Board, int) {
	var winningBoard Board
	for _, number := range numbers {
		for b, board := range boards {
			updated := updateBoard(board, number)
			if boardWins(updated) {
				return updated, number
			}
			boards[b] = updated
		}
	}
	return winningBoard, 0
}

func getLastWinningBoard(boards []Board, numbers []int) (Board, int) {
	var winningBoard Board
	var winningNumber int
	for _, number := range numbers {
		for _, board := range boards {
			// If it's already won, ignore it:
			if boardWins(board) {
				continue
			}
			updated := updateBoard(board, number)
			if boardWins(updated) {
				winningBoard = updated
				winningNumber = number
			}
		}
	}

	return winningBoard, winningNumber
}

func updateBoard(board Board, number int) Board {
	for r, row := range board.rows {
		for s, square := range row {
			if square.value == number {
				square.crossed = true
				row[s] = square
				break
			}
		}
		board.rows[r] = row
	}

	return board
}

func boardWins(board Board) bool {
	winner := false
	// check rows
	for _, row := range board.rows {
		winner = true
		for _, square := range row {
			if square.crossed == false {
				winner = false
			}
		}
		if winner {
			return true
		}
	}
	// check cols
	for i := 0; i < len(board.rows[0]); i++ {
		winner = true
		for _, row := range board.rows {
			square := row[i]
			if square.crossed == false {
				winner = false
			}
		}
		if winner {
			return true
		}
	}

	return winner
}

func calculateBoardScore(board Board, calledNumber int) int {
	sum := 0
	for _, row := range board.rows {
		for _, square := range row {
			if square.crossed == false {
				sum += square.value
			}
		}
	}

	return calledNumber * sum
}
