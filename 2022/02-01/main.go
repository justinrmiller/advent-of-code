package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func computeScore(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return -1, err
	}

	scanner := bufio.NewScanner(f)

	//elf := 1
	score := 0

	for scanner.Scan() {
		//if scanner.Text() == "" {
		//	if elfCalorieCounter > biggestElfCalories {
		//		biggestElf = elf
		//		biggestElfCalories = elfCalorieCounter
		//	}
		//	elf += 1
		//	elfCalorieCounter = 0
		//} else {
		//	calorie, err := strconv.Atoi(scanner.Text())
		//	if err != nil {
		//		return nil, err
		//	}
		//
		//	elfCalorieCounter += calorie
		//
		//	fmt.Printf(
		//		"elf, calorie, elf_calories_counter, biggest_elf: %d, %d, %d, %d\n",
		//		elf,
		//		calorie,
		//		elfCalorieCounter,
		//		biggestElf,
		//	)
		//}
		entry := strings.Split(scanner.Text(), " ")
		opponent := entry[0]
		yours := entry[1]
		shapeComponent := 0
		outcomeComponent := 0

		fmt.Printf("opponent yours: %s %s\n", opponent, yours)

		//shape component calculation
		if yours == "X" {
			shapeComponent = 1
		} else if yours == "Y" {
			shapeComponent = 2
		} else if yours == "Z" {
			shapeComponent = 3
		}

		//outcome component calculation
		if opponent == "A" && yours == "X" {
			fmt.Println("Round ended in a draw.")
			outcomeComponent = 3
		}
		if opponent == "B" && yours == "Y" {
			fmt.Println("Round ended in a draw.")
			outcomeComponent = 3
		}
		if opponent == "C" && yours == "Z" {
			fmt.Println("Round ended in a draw.")
			outcomeComponent = 3
		}

		if opponent == "A" && yours == "Y" {
			outcomeComponent = 6
		}
		if opponent == "A" && yours == "Z" {
			outcomeComponent = 0
		}
		if opponent == "B" && yours == "X" {
			outcomeComponent = 0
		}
		if opponent == "B" && yours == "Z" {
			outcomeComponent = 6
		}
		if opponent == "C" && yours == "X" {
			outcomeComponent = 6
		}
		if opponent == "C" && yours == "Y" {
			outcomeComponent = 0
		}

		fmt.Printf(
			"shape component: %d / outcome component: %d / total: %d\n",
			shapeComponent,
			outcomeComponent,
			shapeComponent+outcomeComponent,
		)
		score = score + shapeComponent + outcomeComponent
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
