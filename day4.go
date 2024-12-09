package main

import (
	"fmt"
	"strings"
)

func FourOne() {
	lines := getLines("inputs/day4test.txt")
	fmt.Println(lines)
	for _, line := range lines {
		for i := range line {
			if line[i] == "X" {

			}
		}
	}
}

func getLines(path string) [][]string {
	scanner, file := GetFileScanner(path)
	defer file.Close()
	var lines [][]string
	for scanner.Scan() {
		line := scanner.Text()
		letters := strings.Split(line, "")
		lines = append(lines, letters)
	}
	return lines
}

func checkEast(lines [][]string, x, y int) bool {
	m := lines[x][y+1]
	a := lines[x][y+2]
	s := lines[x][y+3]
	return m+a+s == "MAS"
}

func checkWest(lines [][]string, x, y int) bool {
	m := lines[x][y-1]
	a := lines[x][y-2]
	s := lines[x][y-3]
	return m+a+s == "MAS"
}

func checkNorth(lines [][]string, x, y int) bool {
	m := lines[x+1][y]
	a := lines[x+2][y]
	s := lines[x+3][y]
	return m+a+s == "MAS"
}

func checkSouth(lines [][]string, x, y int) bool {
	m := lines[x-1][y]
	a := lines[x-2][y]
	s := lines[x-3][y]
	return m+a+s == "MAS"
}
