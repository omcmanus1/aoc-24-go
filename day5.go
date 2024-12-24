package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FiveOne() {
	pairs, updates := parseInput("inputs/day5input.txt")
	verifiedUpdates := verifyUpdates(updates, pairs)
	middleNums := getMiddleNums(verifiedUpdates)
	fmt.Println(middleNums)
}

func getMiddleNums(updates [][]int) int {
	var total int
	for _, update := range updates {
		updateLength := len(update)
		midPoint := int(updateLength / 2)
		middleItem := update[midPoint]
		total += middleItem
	}
	return total
}

func verifyUpdates(updates, pairs [][]int) [][]int {
	var verifiedUpdates [][]int
	for _, update := range updates {
		verified := true
		for index := range update {
			if index > len(update)-2 {
				break
			}
			toCheck := []int{update[index], update[index+1]}
			if !ArrayContainsDeep(pairs, toCheck) {
				verified = false
				break
			}
		}
		if verified {
			verifiedUpdates = append(verifiedUpdates, update)
		}
	}
	return verifiedUpdates
}

func parseInput(path string) ([][]int, [][]int) {
	scanner, file := GetFileScanner(path)
	defer file.Close()

	var pairs [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		pair := strings.Split(line, "|")
		firstNum, _ := strconv.Atoi(pair[0])
		secondNum, _ := strconv.Atoi(pair[1])
		pairs = append(pairs, []int{firstNum, secondNum})
	}

	var updates [][]int
	for scanner.Scan() {
		line := scanner.Text()
		update := strings.Split(line, ",")
		updateInts := ArrayStringToInt(update)
		updates = append(updates, updateInts)
	}

	return pairs, updates
}
