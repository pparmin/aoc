package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	horizontal int
	depth      int
	aim        int
}

func newPosition() *position {
	p := position{horizontal: 0, depth: 0, aim: 0}
	return &p
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
	checkError(scanner.Err())
	return text
}

func partOne() {
	input := ReadFile("input.txt")
	pos := newPosition()

	for _, line := range input {
		input := strings.Split(line, " ")
		command := input[0]
		value, err := strconv.Atoi(input[1])
		checkError(err)

		switch command {
		case "forward":
			pos.horizontal += value
		case "down":
			pos.depth += value
		case "up":
			pos.depth -= value
		}
	}
	fmt.Println("Horizontal position: ", pos.horizontal)
	fmt.Println("Depth: ", pos.depth)
	fmt.Println("Both positions multiplied: ", pos.horizontal*pos.depth)
}

func partTwo() {
	file, err := os.Open("input.txt")
	checkError(err)
	scanner := bufio.NewScanner(file)

	pos := newPosition()

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		command := input[0]
		value, err := strconv.Atoi(input[1])
		checkError(err)

		switch command {
		case "forward":
			pos.horizontal += value
			pos.depth += value * pos.aim
		case "down":
			pos.aim += value
		case "up":
			pos.aim -= value
		}
	}
	checkError(scanner.Err())
	fmt.Println("Horizontal position: ", pos.horizontal)
	fmt.Println("Depth: ", pos.depth)
	fmt.Println("Both positions multiplied: ", pos.horizontal*pos.depth)
}

func main() {
	fmt.Println("PART ONE")
	partOne()
	fmt.Println()
	fmt.Println("PART TWO")
	partTwo()
}
