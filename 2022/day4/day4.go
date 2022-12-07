package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"strconv"
	"strings"
)

// Fill the range
func fillRange(list []string) []int {
	from, _ := strconv.Atoi(list[0])
	to, _ := strconv.Atoi(list[1])
	sectionRange := make([]int, 0)
	for i := from; i <= to; i++ {
		sectionRange = append(sectionRange, i)
	}
	return sectionRange
}

// Check for overlay
func checkOverlay(route []string) bool {
	elf1 := fillRange(strings.Split(route[0], "-"))
	elf2 := fillRange(strings.Split(route[1], "-"))

	left, right := lo.Difference(elf1, elf2)
	return len(left) == 0 || len(right) == 0
}

// Check for partial overlay
func partialOverlay(route []string) bool {
	elf1 := fillRange(strings.Split(route[0], "-"))
	elf2 := fillRange(strings.Split(route[1], "-"))

	return lo.Some(elf1, elf2)
}

func main() {
	part1()
	part2()
}

func part1()  {
	totOverlay := 0
	readFile, _ := os.Open("2022/day4/input")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		route := strings.Split(fileScanner.Text(), ",")
		if checkOverlay(route) {
			totOverlay++
		}
	}
	fmt.Println("Total overlay", totOverlay)
}

func part2()  {
	totPartialOverlay := 0
	readFile, _ := os.Open("2022/day4/input")
	filesScanner := bufio.NewScanner(readFile)
	filesScanner.Split(bufio.ScanLines)

	for filesScanner.Scan() {
		route := strings.Split(filesScanner.Text(), ",")
		if partialOverlay(route) {
			totPartialOverlay++
		}
	}
	fmt.Println("Partial Overlay", totPartialOverlay)
}

