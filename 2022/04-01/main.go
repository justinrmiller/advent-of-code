package main

import (
	"bufio"
	"fmt"
	"os"
)

func subset(first, second []int) bool {
	set := make(map[int]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func computeScore(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return -1, err
	}

	scanner := bufio.NewScanner(f)

	//elf := 1
	score := 0

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		var min1 int
		var max1 int
		var min2 int
		var max2 int
		fmt.Sscanf(text, "%d-%d,%d-%d", &min1, &max1, &min2, &max2)
		fmt.Println(max1)
		range1 := makeRange(min1, max1)
		range2 := makeRange(min2, max2)
		fmt.Println(range1)
		fmt.Println(range2)
		if subset(range1, range2) || subset(range2, range1) {
			score += 1
		}
	}

	err = f.Close()
	if err != nil {
		return -1, err
	}

	return score, nil
}

func main() {
	result, err := computeScore("input.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Printf("Score: %d\n", result)
}
