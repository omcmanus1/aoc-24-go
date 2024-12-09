package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func TwoOne() {
	reports := getReports("day2input.txt")
	safeCount := 0
	for _, report := range reports {
		safe := compareLevels(report)
		safeCount += safe
	}
	fmt.Println("Safe Count: ", safeCount)
}

func TwoTwo() {
	reports := getReports("day2input.txt")
	safeCount := 0
	for _, report := range reports {
		safe := compareLevelsWithDampener(report)
		safeCount += safe
	}
	fmt.Println("Safe Count: ", safeCount)
}

func getReports(path string) [][]int {
	scanner, file := GetFileScanner(path)
	defer file.Close()
	var reports [][]int
	for scanner.Scan() {
		var reportNums []int
		line := scanner.Text()
		reportStrings := strings.Split(line, " ")
		for _, val := range reportStrings {
			valNum, _ := strconv.Atoi(val)
			reportNums = append(reportNums, valNum)
		}
		reports = append(reports, reportNums)
	}
	return reports
}

func compareLevels(report []int) int {
	initialDiff := report[1] - report[0]
	safe := 1
	for i := 0; i < len(report)-1; i++ {
		difference := report[i+1] - report[i]
		if difference*initialDiff <= 0 || int(math.Abs(float64(difference))) > 3 {
			safe = 0
			break
		}
	}
	return safe
}

func compareLevelsWithDampener(report []int) int {
	safe := compareLevels(report)
	if safe == 0 {
		for i := range report {
			reportCopy := arrayCopy(report)
			newReport := append(reportCopy[:i], reportCopy[i+1:]...)
			oneRemovedSafe := compareLevels(newReport)
			if oneRemovedSafe == 1 {
				return 1
			}
		}
	}
	return safe
}
