package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func makeSum(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func partOne() {
	input := ReadFile("input.txt")
	var (
		increases  int = 0
		prevNumber int = 0
	)

	for _, line := range input {
		curNumber, err := strconv.Atoi(line)
		checkError(err)
		if prevNumber < curNumber && prevNumber != 0 {
			increases++
		}
		prevNumber = curNumber
	}
	fmt.Println("Number of increases: ", increases)
}

func partTwo() {
	input := ReadFile("input.txt")
	var (
		values    []int
		sums      []int
		start     int = 0
		end       int = 3
		prevSum   int = 0
		curSum    int = 0
		increases int = 0
	)

	for _, line := range input {
		number, err := strconv.Atoi(line)
		checkError(err)
		values = append(values, number)
	}

	for i := 0; i < len(values); i++ {
		if end > len(values) {
			break
		}

		curSum = makeSum(values[start:end])
		sums = append(sums, curSum)

		if prevSum < curSum && prevSum != 0 {
			increases++
			prevSum = curSum
		}

		prevSum = curSum
		end++
		start++
	}
	fmt.Println("No of increases: ", increases)
}

func main() {
	partOne()
	partTwo()
}
