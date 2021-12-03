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
	resutlTwoOne := handleTwoOne()
	resutlTwoTwo := handleTwoTwo()
	fmt.Println(resultOne)
	fmt.Println(resultTwo)
	fmt.Println(resutlTwoOne)
	fmt.Println(resutlTwoTwo)
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

func handleTwoOne() int {
	inputArr := ReadFile("input2_1.txt")

	horizontalP, dethP := 0, 0

	for _, item := range inputArr {
		itemValues := strings.Split(item, " ")
		itemQuantity, _ := strconv.Atoi(itemValues[1])
		switch itemValues[0] {
		case "forward":
			horizontalP += itemQuantity
		case "up":
			dethP -= itemQuantity
		case "down":
			dethP += itemQuantity
		}
	}

	return horizontalP * dethP
}

func handleTwoTwo() int {
	inputArr := ReadFile("input2_1.txt")

	horizontalP, dethP, aim := 0, 0, 0

	for _, item := range inputArr {
		itemValues := strings.Split(item, " ")
		itemQuantity, _ := strconv.Atoi(itemValues[1])
		switch itemValues[0] {
		case "forward":
			horizontalP += itemQuantity
			dethP += itemQuantity * aim
		case "up":
			aim -= itemQuantity
		case "down":
			aim += itemQuantity
		}
	}

	return horizontalP * dethP
}