package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fileName := "input.txt"

	fmt.Println(part1(fileName))
	fmt.Println(part2(fileName))
}

func part1(fileName string) (max int) {
	file, error := os.Open(fileName)
	defer file.Close()
	if error != nil {
		panic(error)
	}

	scanner := bufio.NewScanner(file)

	maxCalories := 0
	currentElfCalories := 0

	for scanner.Scan() {
		itemCalories, error := strconv.Atoi(scanner.Text())
		currentElfCalories += itemCalories // relies on `AtoI` returning 0 when conversion fails (which it does for empty lines in our input)

		doneCountingForCurrentElf := error != nil
		if doneCountingForCurrentElf {
			if currentElfCalories > maxCalories {
				maxCalories = currentElfCalories
			}
			currentElfCalories = 0 // start counting calories for next elf
		}

	}

	return maxCalories
}

func part2(fileName string) (max int) {
	file, error := os.Open(fileName)
	defer file.Close()
	if error != nil {
		panic(error)
	}

	scanner := bufio.NewScanner(file)

	top3Calories := []int{0, 0, 0}
	currentElfCalories := 0

	for scanner.Scan() {
		itemCalories, error := strconv.Atoi(scanner.Text())
		currentElfCalories += itemCalories // relies on `AtoI` returning 0 when conversion fails (which it does for empty lines in our input)

		doneCountingForCurrentElf := error != nil
		if doneCountingForCurrentElf {
			if currentElfCalories > top3Calories[0] {
				top3Calories[0] = currentElfCalories
				sort.Ints(top3Calories)
			}
			currentElfCalories = 0 // start counting calories for next elf
		}

	}

	return top3Calories[0] + top3Calories[1] + top3Calories[2]
}
