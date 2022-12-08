package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Elf struct {
	number   int
	calories int
}

func newElf(number int, cal int) *Elf {
	e := Elf{
		number:   number,
		calories: cal,
	}
	return &e
}

func readFile(input string) []string {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ProcessInput(input string) int {
	lines := readFile(input)

	var elves []Elf
	var calories int
	elfCounter := 1

	for _, line := range lines {
		if len(line) == 0 {
			elf := newElf(elfCounter, calories)
			elves = append(elves, *elf)
			elfCounter++
			calories = 0
			continue
		}
		v, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		calories += v
	}

	highestCal := 0
	for _, elf := range elves {
		// fmt.Printf("elf n#%d --> calories: %d\n", elf.number, elf.calories)
		if elf.calories > highestCal {
			highestCal = elf.calories
		}
	}
	return highestCal
}

func main() {
	mostCalories := ProcessInput("input.txt")
	fmt.Printf("Most calories carried: %d", mostCalories)
}

//
