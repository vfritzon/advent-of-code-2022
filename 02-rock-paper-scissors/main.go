package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// A, X = rock, 1
	// B, Y = paper, 2
	// C, Z = scissors, 3
	scoreMap1 := map[string]int{
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

	// A = rock, 1
	// B = paper, 2
	// C = scissors, 3
	// X = lose
	// Y = draw
	// Z = win
	scoreMap2 := map[string]int{
		"A X": 3 + 0,
		"A Y": 1 + 3,
		"A Z": 2 + 6,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 2 + 0,
		"C Y": 3 + 3,
		"C Z": 1 + 6,
	}
	fmt.Println(calculateScore("input.txt", scoreMap1))
	fmt.Println(calculateScore("input.txt", scoreMap2))
}

func calculateScore(fileName string, scoreMap map[string]int) int {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	totalScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		totalScore += scoreMap[scanner.Text()]
	}
	return totalScore
}
