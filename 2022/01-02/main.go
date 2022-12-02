package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func extractBiggestElf(filename string) ([]int, error) {
	calorieCounts := make([]int, 10)

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	elf := 1
	elfCalorieCounter := 0

	for scanner.Scan() {
		if scanner.Text() == "" || scanner.Text() == "-" {
			calorieCounts = append(calorieCounts, elfCalorieCounter)
			elf += 1
			elfCalorieCounter = 0
		} else {
			calorie, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return nil, err
			}

			elfCalorieCounter += calorie

			//fmt.Printf(
			//	"elf, calorie, elf_calories_counter: %d, %d, %d\n",
			//	elf,
			//	calorie,
			//	elfCalorieCounter,
			//)
		}
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	return calorieCounts, nil
}

/*
	Note: CSV file needs a - delimiter

Example:
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000

-
*/

func main() {
	calorieCounts, err := extractBiggestElf("input.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	sort.Slice(calorieCounts, func(i, j int) bool {
		return calorieCounts[i] > calorieCounts[j]
	})
	sum := 0
	for i := 0; i < 3; i++ {
		fmt.Println(calorieCounts[i])
		sum += calorieCounts[i]
	}
	fmt.Printf("Sum of Calorie Counts, Number of Calorie Counts: %d, %d\n", sum, len(calorieCounts))
}
