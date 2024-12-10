package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func GetFileScanner(path string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("file open error:", err)
	}
	scanner := bufio.NewScanner(file)
	return scanner, file
}

func GetKey(m map[int]int) int {
	for k := range m {
		return k
	}
	return 0
}

func GetValue(m map[int]int) int {
	for _, v := range m {
		return v
	}
	return 0
}

func StringToInt(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return output
}

func arrayCopy(orig []int) []int {
	newNums := make([]int, len(orig))
	copy(newNums, orig[:])
	return newNums
}

func ArrayContainsDeep(coords [][]int, target []int) bool {
	for _, c := range coords {
		if reflect.DeepEqual(c, target) {
			return true
		}
	}
	return false
}
