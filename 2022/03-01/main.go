package main

import (
	"bufio"
	"fmt"
	"github.com/juliangruber/go-intersect"
	"os"
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
		text := scanner.Text()
		lhs := text[0 : len(text)/2+1]
		rhs := text[len(text)/2+1 : len(text)]
		fmt.Printf("text lhs rhs: %s %s %s\n", text, lhs, rhs)
		commonarr := intersect.Simple(lhs, rhs)
		if len(commonarr) == 0 {
			// skip
		} else {
			fmt.Println(commonarr)
			common := string(commonarr[0].(uint8))
			code := uint8(0)
			//if common <= 122 && common >= 97 {
			//	code = common - 96
			//} else {
			//	code = common - 38
			//}
			fmt.Printf("character - code: %s - %d\n", common, code)

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
