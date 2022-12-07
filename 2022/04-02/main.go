package main

import (
	"bufio"
	"fmt"
	"os"
)

func intersection(first, second []int) []int {
	out := []int{}
	bucket := map[int]bool{}

	for _, i := range first {
		for _, j := range second {
			if i == j && !bucket[i] {
				out = append(out, i)
				bucket[i] = true
			}
		}
	}

	return out
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
		if len(intersection(range1, range2)) != 0 {
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
