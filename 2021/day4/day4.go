package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	value  string
	marked bool
}

type Row struct {
	values     []Number
	horizontal int
}

type Board struct {
	rows []Row
	// holds a counter for each row; if one row reaches 5, we know it's full
	horizontal int
	// holds a counter for each position in each row; if one pos reaches 5
	// we know the column is full
	vertical map[int]int

	// only used in part two
	ready bool
}

func newRow() *Row {
	r := Row{
		values:     make([]Number, 0),
		horizontal: 0,
	}
	return &r
}

func newBoard() *Board {
	b := Board{
		rows:       make([]Row, 0),
		horizontal: 0,
		vertical: map[int]int{
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
		},
	}
	return &b
}

func getBoards(input []string) []Board {
	b := newBoard()
	row := newRow()
	var (
		boards   []Board
		values   []string
		rowCount int = 0
	)

	for i, line := range input {
		if i == 0 {
			continue
		}
		if line == "" {
			continue
		}
		if rowCount > 4 {
			boards = append(boards, *b)
			rowCount = 0
			b = newBoard()
		}

		values = strings.Split(line, " ")
		for _, v := range values {
			if v == "" {
				continue
			} else {
				n := Number{
					value:  v,
					marked: false,
				}
				row.values = append(row.values, n)
			}
		}
		b.rows = append(b.rows, *row)
		row = newRow()
		rowCount++
	}
	return boards
}

func play(n string, boards []Board) (Board, bool) {
	for _, b := range boards {
		for k, r := range b.rows {
			for l, num := range r.values {
				if num.value == n {
					num.marked = true
					r.horizontal++
					b.vertical[l] += 1
					b.rows[k].values[l] = num
					b.rows[k].horizontal = r.horizontal
				}
			}
		}
		if isComplete(b) {
			b.ready = true
			return b, true
		}
	}
	return Board{}, false
}

func playTwo(n string, b Board) Board {
	// stop execution for boards which are already ready
	if b.ready {
		return Board{}
	}

	for k, r := range b.rows {
		for l, num := range r.values {
			if num.value == n {
				num.marked = true
				r.horizontal++
				b.vertical[l] += 1
				b.rows[k].values[l] = num
				b.rows[k].horizontal = r.horizontal
			}
		}
	}
	if isComplete(b) {
		b.ready = true
		return b
	}
	return Board{}
}

func isComplete(b Board) bool {
	for _, v := range b.vertical {
		if v == 5 {
			return true
		}
	}

	for _, row := range b.rows {
		if row.horizontal == 5 {
			return true
		}
	}
	return false
}

func calcScore(b Board) int {
	sum := 0
	for _, row := range b.rows {
		for _, v := range row.values {
			if !v.marked {
				val, err := strconv.Atoi(v.value)
				checkError(err)
				sum += val
			}
		}
	}
	return sum
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(name string) []string {
	file, err := os.Open(name)
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text []string
	for scanner.Scan() {
		line := scanner.Text()
		text = append(text, line)
	}
	return text
}

func partOne() {
	input := ReadFile("input.txt")
	complete := false
	var final Board

	numbers := strings.Split(input[0], ",")
	boards := getBoards(input)

	for _, n := range numbers {
		final, complete = play(n, boards)
		if complete {
			fmt.Println("FINAL BOARD:", final)
			result := calcScore(final)
			lastNum, err := strconv.Atoi(n)
			checkError(err)
			fmt.Println("FINAL RESULT: ", result*lastNum)
			return
		}
	}
}

func partTwo() {
	input := ReadFile("input.txt")
	var completed []Board
	var r Board
	result := 0
	numbers := strings.Split(input[0], ",")
	boards := getBoards(input)

	for _, n := range numbers {
		for i, b := range boards {
			r = playTwo(n, b)
			if r.ready {
				completed = append(completed, r)
				if len(completed) == len(boards) {
					fmt.Println("Last board to be ready reached!")
					fmt.Println("Board no ", i, ": ", r)
					result = calcScore(r)
					lastNum, err := strconv.Atoi(n)
					checkError(err)
					fmt.Println("FINAL RESULT: ", result*lastNum)
					return
				}
				// replace old board with new board in overall boards slice
				boards[i] = r
			}
		}
	}
}

func main() {
	partOne()
	partTwo()
}
