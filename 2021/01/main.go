package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func readInput(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	var depths []int

	for _, line := range lines {
		depth, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, err
		}
		depths = append(depths, depth)
	}

	return depths, nil
}

func main() {
	depths, err := readInput("input.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	count := 0

	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			count++
		}
		// fmt.Printf("count %d, current depth: %d\n", count, depths[i])
	}
	fmt.Printf("Count of depth increases: %d\n", count)

	var sums []int
	for i := 0; i <= len(depths)-3; i++ {
		first, second, third := depths[i], depths[i+1], depths[i+2]
		sum := first + second + third
		sums = append(sums, sum)
	}
	sumGreaterCount := 0
	for i := 1; i < len(sums); i++ {
		if sums[i] > sums[i-1] {
			sumGreaterCount++
		}
	}
	fmt.Printf("Sum greater than previous count: %d\n", sumGreaterCount)
}
