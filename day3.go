package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ThreeOne() {
	text := getFileText("inputs/day3input.txt")
	total := addMultiplications(text)
	fmt.Println("Total: ", total)
}

func ThreeTwo() {
	text := getFileText("inputs/day3input.txt")
	groups := strings.Split(text, "don't()")
	var matches string
	for i, group := range groups {
		if i == 0 {
			matches += group
		} else {
			_, after, _ := strings.Cut(group, "do()")
			if len(after) > 1 {
				matches += after
			}
		}
	}
	total := addMultiplications(matches)
	fmt.Println("Total: ", total)
}

func getFileText(path string) string {
	scanner, file := GetFileScanner(path)
	defer file.Close()
	var fileText string
	for scanner.Scan() {
		line := scanner.Text()
		fileText += line
	}
	return fileText
}

func addMultiplications(text string) int {
	mulRe := regexp.MustCompile(`mul\(\d+,\d+\)`)
	numsRe := regexp.MustCompile(`(\d+),(\d+)`)
	total := 0
	matches := mulRe.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		first := numsRe.FindAllStringSubmatch(match[0], -1)[0][1]
		firstNum, _ := strconv.Atoi(first)
		second := numsRe.FindAllStringSubmatch(match[0], -1)[0][2]
		secondNum, _ := strconv.Atoi(second)
		total += firstNum * secondNum
	}
	return total
}
