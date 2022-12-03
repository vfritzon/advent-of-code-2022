package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	roundScore := map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}

	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	totalScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		totalScore += roundScore[scanner.Text()]
	}

	fmt.Print(totalScore)
}
