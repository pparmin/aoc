package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type report struct {
	gamma     int64
	epsilon   int64
	oxRating  int64
	co2Rating int64
}

func newReport() *report {
	r := report{gamma: 0, epsilon: 0, oxRating: 0, co2Rating: 0}
	return &r
}

func calcBits(binaries []string, i int) (int, int) {
	z, o := 0, 0
	for _, b := range binaries {
		if b[i] == '0' {
			z++
		} else if b[i] == '1' {
			o++
		}
	}
	return z, o
}

func calculateRating(binaries []string, criteria string) string {
	var rating string
	i := 0
	zeroes, ones := 0, 0
	for i < 12 {
		zeroes, ones = calcBits(binaries, i)
		fmt.Printf("ZEROES: %d, ONES: %d\n", zeroes, ones)

		switch criteria {
		case "most":
			if zeroes > ones {
				binaries = remove(binaries, i, '1')
				fmt.Println("MOST COMMON -- MORE ZEROES THAN ONES FOUND -- REMOVING all 1s at index: ", i)
			} else if zeroes < ones {
				binaries = remove(binaries, i, '0')
				fmt.Println("MOST COMMON -- MORE ONES THAN ZEROES -- REMOVING all 0s at index: ", i)
			} else if zeroes == ones {
				binaries = remove(binaries, i, '0')
				fmt.Println("MOST COMMON -- EQUAL ONES AND ZEROES -- REMOVING all 1s at index: ", i)
			}

		case "least":
			if zeroes > ones {
				binaries = remove(binaries, i, '0')
				fmt.Println("LEAST COMMON -- MORE ZEROES THAN ONES FOUND -- REMOVING all 0s at index: ", i)
			} else if zeroes < ones {
				binaries = remove(binaries, i, '1')
				fmt.Println("LEAST COMMON -- MORE ONES THAN ZEROES -- REMOVING all 1s at index: ", i)
			} else if zeroes == ones {
				binaries = remove(binaries, i, '1')
				fmt.Println("LEAST COMMON --  EQUAL ONES AND ZEROES -- REMOVING all 0s at index: ", i)
			}
		}
		if len(binaries) == 1 {
			rating = binaries[0]
			fmt.Println("FOUND LAST NUMBER: ", rating)
			fmt.Println()
			break
		}
		zeroes, ones = 0, 0
		i++
		fmt.Println("NEW LIST: ", binaries)
	}
	return rating
}

func remove(binary []string, index int, bit byte) []string {
	new_list := make([]string, 0)
	for _, b := range binary {
		if b[index] != bit {
			new_list = append(new_list, b)
		}
	}
	return new_list
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
	var err error
	var (
		gammaBuilder       = strings.Builder{}
		epsilonBuilder     = strings.Builder{}
		r                  = newReport()
		zeroes         int = 0
		ones           int = 0
		binaries       []string
	)

	for _, line := range input {
		binary := line
		binaries = append(binaries, binary)
	}

	i := 0
	for i < 12 {
		zeroes, ones = calcBits(binaries, i)
		fmt.Printf("ITER %d: zeroes %d, ones %d\n", i, zeroes, ones)

		if zeroes > ones {
			gammaBuilder.WriteRune('0')
			epsilonBuilder.WriteRune('1')

		} else if zeroes < ones {
			gammaBuilder.WriteRune('1')
			epsilonBuilder.WriteRune('0')
		}

		fmt.Printf("ITER %d: gamma binary: %q\n", i, gammaBuilder.String())
		fmt.Printf("ITER %d: epsilon binary: %q\n", i, epsilonBuilder.String())
		r.gamma, err = strconv.ParseInt(gammaBuilder.String(), 2, 16)
		checkError(err)

		r.epsilon, err = strconv.ParseInt(epsilonBuilder.String(), 2, 16)
		checkError(err)

		fmt.Printf("ITER %d: gamma decimal: %d\n", i, r.gamma)
		fmt.Printf("ITER %d: epsilon decimal: %d\n", i, r.epsilon)

		zeroes, ones = 0, 0
		i++
		fmt.Println()
	}
	fmt.Println("Power consumption: ", r.gamma*r.epsilon)
	fmt.Println()
}

func partTwo() {
	input := ReadFile("input.txt")
	var err error
	var (
		oxygen []string
		co2    []string
		r      = newReport()
	)

	for _, line := range input {
		binary := line
		oxygen = append(oxygen, binary)
		co2 = append(co2, binary)
	}

	oxBinary := calculateRating(oxygen, "most")
	co2Binary := calculateRating(co2, "least")

	r.oxRating, err = strconv.ParseInt(oxBinary, 2, 16)
	checkError(err)
	r.co2Rating, err = strconv.ParseInt(co2Binary, 2, 16)
	checkError(err)
	fmt.Printf("IN BYTES: OX: %q, CO2: %q\n", oxBinary, co2Binary)
	fmt.Printf("DECIMAL: OX: %d, CO2: %d\n", r.oxRating, r.co2Rating)
	fmt.Printf("Life support rating: %d\n", r.oxRating*r.co2Rating)
}

func main() {
	partOne()
	partTwo()
}
