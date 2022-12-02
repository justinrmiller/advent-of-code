package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Instruction struct {
	command string
	amount  int
}

func readInput(filename string) ([]Instruction, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(f)
	reader.Comma = ' '

	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	var instructions []Instruction

	for _, line := range lines {
		command := line[0]
		amount, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}
		instruction := Instruction{command: command, amount: amount}
		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

func main() {
	instructions, err := readInput("input.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	horizontal := 0
	depth := 0

	for i := 0; i < len(instructions); i++ {
		//fmt.Printf("% +v\n", instructions[i])
		switch instructions[i].command {
		case "forward":
			horizontal += instructions[i].amount
		case "up":
			depth -= instructions[i].amount
		case "down":
			depth += instructions[i].amount
		}
	}
	fmt.Printf("Part 1: Count of instructions: %d\n", len(instructions))
	fmt.Printf("Part 2: Depth: %d\nHorizontal: %d\n", depth, horizontal)
	fmt.Printf("Part 3: Depth * Horizontal: %d\n", depth*horizontal)

	horizontal = 0
	depth = 0
	aim := 0

	for i := 0; i < len(instructions); i++ {
		//fmt.Printf("% +v\n", instructions[i])
		switch instructions[i].command {
		case "forward":
			horizontal += instructions[i].amount
			depth += aim * instructions[i].amount
		case "up":
			aim -= instructions[i].amount
		case "down":
			aim += instructions[i].amount
		}
	}
	fmt.Printf("Part 2: Count of instructions: %d\n", len(instructions))
	fmt.Printf("Part 2: Depth: %d\nHorizontal: %d\n", depth, horizontal)
	fmt.Printf("Part 2: Depth * Horizontal: %d\n", depth*horizontal)
}
