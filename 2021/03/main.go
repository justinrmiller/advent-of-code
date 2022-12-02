package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func readInput(filename string) ([]string, error) {
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

	var bitsArray []string

	for _, line := range lines {
		bits := line[0]
		bitsArray = append(bitsArray, bits)
	}

	return bitsArray, nil
}

func main() {
	bitsArray, err := readInput("input.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	//count := 0
	var gammaTracker [12]int
	var epsilonTracker [12]int

	for i := 0; i < len(bitsArray); i++ {
		fmt.Printf("position %d, current bit: %s\n", i, bitsArray[i])
		for j := 0; j < len(bitsArray[i]); j++ {
			if bitsArray[i][j] == '0' {
				gammaTracker[j] += 1
			} else {
				gammaTracker[j] -= 1
			}
		}
	}

	for i := 0; i < len(bitsArray); i++ {
		fmt.Printf("position %d, current bit: %s\n", i, bitsArray[i])
		for j := 0; j < len(bitsArray[i]); j++ {
			if bitsArray[i][j] == '1' {
				epsilonTracker[j] += 1
			} else {
				epsilonTracker[j] -= 1
			}
		}
	}

	var gammaStringRepr string
	for i := 0; i < len(gammaTracker); i++ {
		if gammaTracker[i] >= 0 {
			fmt.Printf("position %d, gamma tracker: %d, signal: 0\n", i, gammaTracker[i])
			gammaStringRepr += "0"
		} else {
			fmt.Printf("position %d, gamma tracker: %d, signal: 1\n", i, gammaTracker[i])
			gammaStringRepr += "1"
		}
	}
	gammaIntRepr, err := strconv.ParseInt(gammaStringRepr, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("int repr for gamma: %d\n", gammaIntRepr)

	var epsilonStringRepr string
	for i := 0; i < len(epsilonTracker); i++ {
		if epsilonTracker[i] >= 0 {
			fmt.Printf("position %d, epsilon tracker: %d, signal: 0\n", i, epsilonTracker[i])
			epsilonStringRepr += "0"
		} else {
			fmt.Printf("position %d, epsilon tracker: %d, signal: 1\n", i, epsilonTracker[i])
			epsilonStringRepr += "1"
		}
	}

	epsilonIntRepr, err := strconv.ParseInt(epsilonStringRepr, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("int repr for epsilon: %d\n", epsilonIntRepr)

	fmt.Printf("gamma * epsilon: %d\n", gammaIntRepr*epsilonIntRepr)
}
