package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(name string) []string{
	dat, err := os.ReadFile(fmt.Sprintf("./input/%s", name))
	check(err)
	array := strings.Split(string(dat), "\n")
	return array
}

func main() {
	resultOne := handleOne()
	resultTwo := handleTwo()
	fmt.Println(resultOne)
	fmt.Println(resultTwo)
}

func handleOne() int {
	inputArr := ReadFile("input1.txt")

	count := 0
	for i, item := range inputArr {
		if i == 0 {
			continue
		}

		value, _ := strconv.Atoi(item)
		previous, _ := strconv.Atoi(inputArr[i-1])
		if value > previous {
			count++
		}
	}

	return count
}

func handleTwo() int {
	// read file
	inputArr := ReadFile("input1.txt")

	//sum each 3
	var newArr []int
	for i, item := range inputArr {
		if i > len(inputArr) - 3 {
			break
		}
		value, _ := strconv.Atoi(item)
		valueTwo, _ := strconv.Atoi(inputArr[i+1])
		valueThree, _ := strconv.Atoi(inputArr[i+2])
		newArr = append(newArr, value + valueTwo + valueThree)
	}

	count := 0
	for i, item := range newArr {
		if i == 0 {
			continue
		}

		if item > newArr[i-1] {
			count++
		}
	}

	return count
}

