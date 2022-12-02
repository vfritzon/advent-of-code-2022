package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	maxCalories := 0
	currentElfCalories := 0
	for {
		lastItem := !scanner.Scan()
		line := scanner.Text()
		lastItemForElf := line == ""

		if lastItem || lastItemForElf {
			if currentElfCalories > maxCalories {
				maxCalories = currentElfCalories
			}
			currentElfCalories = 0
		} else {
			itemCalories, _ := strconv.Atoi(line)
			currentElfCalories += itemCalories
		}

		if lastItem {
			break
		}
	}

	fmt.Print(maxCalories)
}
