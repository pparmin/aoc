package day5_test

import (
	d5 "aoc/day5"
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	t.Parallel()
	t.Run("test input part one", func(t *testing.T) {
		input := d5.ReadFile("test.txt")
		got := d5.GetLines(input)
		fmt.Println("GOT: ", got)
		board := d5.Initialize(got)
		fmt.Println("BOARD: ")
		fmt.Println(board)
		_ = d5.CalcOverlap(got, board)
	})

	t.Run("real input part one", func(t *testing.T) {
		// input := d5.ReadFile("input.txt")
		// got := d5.GetLines(input)
		// fmt.Println("GOT: ", got)
	})
}
