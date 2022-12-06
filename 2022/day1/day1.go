package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	caloriesSlice := buildCalorieSlice()
	println(caloriesSlice)
	sort.Ints(caloriesSlice)
	highestCalorie := caloriesSlice[len(caloriesSlice)-1]
	println("highest:", highestCalorie)
}

func part2() {
	caloriesSlice := buildCalorieSlice()
	sort.Ints(caloriesSlice)
	top3 := caloriesSlice[len(caloriesSlice)-3:]
	tot := 0
	for _, v := range top3 {
		tot += v
	}
	println("Sum top3:", tot)
}

func buildCalorieSlice() []int {
	readFile, err := os.Open("2022/day1/input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	calories := 0
	caloriesSlice := make([]int, 0)

	for fileScanner.Scan() {
		var strVar = fileScanner.Text()
		if len(strVar) != 0 {
			intVar, _ := strconv.Atoi(strVar)
			calories += intVar
		} else {
			caloriesSlice = append(caloriesSlice, calories)
			calories = 0
		}
	}
	readFile.Close()
	return caloriesSlice
}
