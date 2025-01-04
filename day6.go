package main

import (
	"fmt"
	"slices"
	"strings"
)

func SixOne() {
	grid, coords := populateGrid("inputs/day6input.txt")
	direction := "up"
	finished := false
	newPositionCount := 0
	for !finished {
		grid, coords = makeMove(grid, coords, &direction, &finished, &newPositionCount)
		printGrid(grid)
	}
	fmt.Println("count: ", newPositionCount)
}

func makeMove(grid [][]string, coords []int, direction *string, finished *bool, count *int) ([][]string, []int) {
	fmt.Println("coords: ", coords)
	fmt.Println("direction: ", *direction)
	nextCoords := []int{}
	if coords[1] > len(grid)-1 || coords[0] < 0 {
		*finished = true
		fmt.Println("finished")
		return grid, nextCoords
	}
	switch *direction {
	case "up":
		if grid[coords[1]][coords[0]] == "#" {
			*direction = "right"
			nextCoords = []int{coords[0], coords[1] + 1}
			return grid, nextCoords
		}
		nextCoords = []int{coords[0], coords[1] - 1}
	case "right":
		if coords[0] > len(grid[0])-1 {
			*finished = true
			return grid, nextCoords
		}
		if grid[coords[1]][coords[0]] == "#" {
			*direction = "down"
			nextCoords = []int{coords[0] - 1, coords[1]}
			return grid, nextCoords
		}
		nextCoords = []int{coords[0] + 1, coords[1]}
	case "down":
		if grid[coords[1]][coords[0]] == "#" {
			*direction = "left"
			nextCoords = []int{coords[0], coords[1] - 1}
			return grid, nextCoords
		}
		nextCoords = []int{coords[0], coords[1] + 1}
	case "left":
		if grid[coords[1]][coords[0]] == "#" {
			*direction = "up"
			nextCoords = []int{coords[0] + 1, coords[1]}
			return grid, nextCoords
		}
		nextCoords = []int{coords[0] - 1, coords[1]}
	}
	if grid[coords[1]][coords[0]] == "^" || grid[coords[1]][coords[0]] == "." {
		grid[coords[1]][coords[0]] = "X"
		*count++
	}
	return grid, nextCoords
}

func populateGrid(path string) ([][]string, []int) {
	scanner, file := GetFileScanner(path)
	defer file.Close()
	var grid [][]string
	var coords []int
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		symbols := strings.Split(line, "")
		startPoint := slices.Index(symbols, "^")
		if startPoint > 0 {
			coords = []int{startPoint, lineNum}
		}
		grid = append(grid, symbols)
		lineNum++
	}
	return grid, coords
}

func printGrid(grid [][]string) {
	for _, line := range grid {
		fmt.Println(line)
	}
}
