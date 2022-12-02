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
	file.Close()

	return maxCalories
}

func part2(fileName string) (max int) {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	top3 := []int{0, 0, 0}
	currentElfCalories := 0
	for {
		lastItem := !scanner.Scan()
		line := scanner.Text()
		lastItemForElf := line == ""

		if lastItem || lastItemForElf {
			if currentElfCalories > top3[0] {
				top3[0] = currentElfCalories
				sort.Ints(top3)
			}
			currentElfCalories = 0
		} else {
			itemCalories, _ := strconv.Atoi(line)
			currentElfCalories += itemCalories
		}

		if lastItem {
			file.Close()
			break
		}
	}

	sum := 0
	for _, v := range top3 {
		sum += v
	}

	return sum
}
