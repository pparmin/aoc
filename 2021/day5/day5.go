package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	X, Y  int
	count int
}

type Line struct {
	start Point
	end   Point
}

func NewPoint() *Point {
	p := Point{X: 0, Y: 0, count: 0}
	return &p
}

func InitPoint(x, y int) *Point {
	p := Point{X: x, Y: y, count: 0}
	return &p
}

func NewLine() *Line {
	l := Line{start: Point{}, end: Point{}}
	return &l
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

func parseNumber(line string) []string {
	var number []rune
	var result []string

	for i, char := range line {
		switch char {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			number = append(number, char)
			// handle end of line
			if i == len(line)-1 {
				result = append(result, string(number))
			}
		case ',', '>':
			result = append(result, string(number))
			number = []rune{}
		case ' ', '-':
			continue
		}
	}
	return result
}

func findMax(lines []Line) (int, int) {
	max_x := 0
	max_y := 0
	for _, l := range lines {
		if l.start.X > max_x {
			max_x = l.start.X
		} else if l.start.Y > max_y {
			max_y = l.start.Y
		} else if l.end.X > max_x {
			max_x = l.end.X
		} else if l.end.Y > max_y {
			max_y = l.end.Y
		}
	}
	return max_x, max_y
}

// parse input to generate the field of hydrothermal vents
func GetLines(input []string) []Line {
	var lines []Line
	var err error
	for _, line := range input {
		// fmt.Println("LINE: ", line)
		l := NewLine()
		p := NewPoint()

		numbers := parseNumber(line)
		// fmt.Println("Numbers: ", numbers)
		p.X, err = strconv.Atoi(numbers[0])
		checkError(err)
		p.Y, err = strconv.Atoi(numbers[1])
		checkError(err)
		// fmt.Println("CURRENT POINT: ", p)
		l.start = *p
		// fmt.Println("CURRENT LINE: ", l)

		p.X, err = strconv.Atoi(numbers[2])
		checkError(err)
		p.Y, err = strconv.Atoi(numbers[3])
		checkError(err)
		// fmt.Println("CURRENT POINT: ", p)
		l.end = *p
		// fmt.Println("CURRENT LINE: ", l)
		lines = append(lines, *l)
	}
	return lines
}

func Initialize(lines []Line) [][]Point {
	var board [][]Point
	var horizontal []Point
	MAX_X, MAX_Y := findMax(lines)
	for i := 0; i <= MAX_Y; i++ {
		for j := 0; j <= MAX_X; j++ {
			p := InitPoint(j, i)
			horizontal = append(horizontal, *p)
		}
		board = append(board, horizontal)
		horizontal = []Point{}
	}
	return board
}

func CalcOverlap(lines []Line, board [][]Point) int {
	//MAX := findMax(lines)
	for _, l := range lines {
		for _, line := range board {
			// ignore lines which are diagonal
			if l.start.X == l.end.X {
				fmt.Println("VERTICAL: ", l)
				fmt.Println("b", line)
			} else if l.start.Y == l.end.Y {
				fmt.Println("HORIZONTAL: ", l)
			} else {
				continue
			}
		}
	}
	return 0
}
