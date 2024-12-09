package main

import (
	"fmt"
	"strings"
)

const word = "XMAS"

func FourOne() {
	lines := getLines("inputs/day4input.txt")
	found := 0
	for x := 0; x < len(lines); x++ {
		for y := 0; y < len(lines[x]); y++ {
			if lines[x][y] == "X" {
				checkEast(lines, x, y, &found)
				checkWest(lines, x, y, &found)
				checkNorth(lines, x, y, &found)
				checkSouth(lines, x, y, &found)
				checkNorthEast(lines, x, y, &found)
				checkNorthWest(lines, x, y, &found)
				checkSouthEast(lines, x, y, &found)
				checkSouthWest(lines, x, y, &found)
			}
		}
	}
	fmt.Printf("Found %d occurrences of %s\n", found, word)
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

func checkEast(lines [][]string, x, y int, found *int) {
	if y+3 < len(lines[x]) {
		for i := 0; i < 4; i++ {
			if lines[x][y+i] != string(word[i]) {
				return
			}
		}
		*found++
	}
}

func checkWest(lines [][]string, x, y int, found *int) {
	if y-3 >= 0 {
		for i := 0; i < 4; i++ {
			if lines[x][y-i] != string(word[i]) {
				return
			}
		}
		*found++
	}
}

func checkNorth(lines [][]string, x, y int, found *int) {
	if x-3 >= 0 {
		for i := 0; i < 4; i++ {
			if lines[x-i][y] != string(word[i]) {
				return
			}
		}
		*found++
	}
}

func checkSouth(lines [][]string, x, y int, found *int) {
	if x+3 < len(lines) {
		for i := 0; i < 4; i++ {
			if lines[x+i][y] != string(word[i]) {
				return
			}
		}
		*found++
	}
}

func checkNorthEast(lines [][]string, x, y int, found *int) {
	if x-3 >= 0 && y+3 < len(lines[x]) {
		for i := 0; i < 4; i++ {
			if lines[x-i][y+i] != string(word[i]) {
				return
			}
		}
		*found++
	}
}

func checkNorthWest(lines [][]string, x, y int, found *int) {
	if x-3 >= 0 && y-3 >= 0 {
		for i := 0; i < 4; i++ {
			if lines[x-i][y-i] != string(word[i]) {
				return
			}
		}
		*found++
	}
}

func checkSouthEast(lines [][]string, x, y int, found *int) {
	if x+3 < len(lines) && y+3 < len(lines[x]) {
		for i := 0; i < 4; i++ {
			if lines[x+i][y+i] != string(word[i]) {
				return
			}
		}
		*found++
	}
}

func checkSouthWest(lines [][]string, x, y int, found *int) {
	if x+3 < len(lines) && y-3 >= 0 {
		for i := 0; i < 4; i++ {
			if lines[x+i][y-i] != string(word[i]) {
				return
			}
		}
		*found++
	}
}
