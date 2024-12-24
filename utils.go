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

func ArrayCopy(orig []int) []int {
	newNums := make([]int, len(orig))
	copy(newNums, orig[:])
	return newNums
}

func ArrayContainsDeep(nums [][]int, target []int) bool {
	for _, num := range nums {
		if reflect.DeepEqual(num, target) {
			return true
		}
	}
	return false
}

func ArrayStringToInt(input []string) []int {
	var output []int
	for _, val := range input {
		intVal, _ := strconv.Atoi(val)
		output = append(output, intVal)
	}
	return output
}
