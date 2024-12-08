package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func OneOne() {
	listOne, listTwo := buildLists("inputs/day1input.txt")
	sortListDescending(listOne)
	sortListDescending(listTwo)
	var totalDifference int
	for i := range listOne {
		difference := int(math.Abs(float64(listOne[i] - listTwo[i])))
		totalDifference += difference
	}
	fmt.Println("Total Difference: ", totalDifference)
}

func OneTwo() {
	listOne, listTwo := buildLists("inputsday1input.txt")
	similarityScore := 0
	for _, listOneVal := range listOne {
		instanceCount := 0
		for _, listTwoVal := range listTwo {
			if listOneVal == listTwoVal {
				instanceCount++
			}
		}
		similarityScore += listOneVal * instanceCount
	}
	fmt.Println("Similarity Score: ", similarityScore)
}

func buildLists(path string) ([]int, []int) {
	scanner, file := GetFileScanner(path)
	defer file.Close()
	var listOne []int
	var listTwo []int
	for scanner.Scan() {
		line := scanner.Text()
		locationIds := strings.Split(line, "   ")
		firstId, _ := strconv.Atoi(locationIds[0])
		secondId, _ := strconv.Atoi(locationIds[1])
		listOne = append(listOne, firstId)
		listTwo = append(listTwo, secondId)
	}
	return listOne, listTwo
}

func sortListDescending(list []int) {
	sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j]
	})
}
