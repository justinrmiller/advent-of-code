package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func extractBiggestElf(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	elf := 1
	elfCalorieCounter := 0
	biggestElfCalories := 0
	biggestElf := -1

	for scanner.Scan() {
		if scanner.Text() == "" {
			if elfCalorieCounter > biggestElfCalories {
				biggestElf = elf
				biggestElfCalories = elfCalorieCounter
			}
			elf += 1
			elfCalorieCounter = 0
		} else {
			calorie, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return nil, err
			}

			elfCalorieCounter += calorie

			fmt.Printf(
				"elf, calorie, elf_calories_counter, biggest_elf: %d, %d, %d, %d\n",
				elf,
				calorie,
				elfCalorieCounter,
				biggestElf,
			)
		}
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	return []int{biggestElf, biggestElfCalories}, nil
}

func main() {
	result, err := extractBiggestElf("input.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Printf("Biggest Elf: %d, Biggest Elf Calories: %d\n", result[0], result[1])
}
