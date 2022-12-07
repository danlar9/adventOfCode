package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"strings"
)

/** Error function **/
func check(e error) {
	if e != nil {
		panic(e)
	}
}

/** Helper function to fill priorities 1:a, ... , Z:52 */
func toCharStr(i int) string {
	if i <= 26 {
		return string(rune('a' - 1 + i))
	} else {
		return string(rune('A' - 27 + i))
	}
}

/** Get a priority map [ a:1, ... , Z:52 ] **/
func getPriorities() map[string]int {
	priorities := make(map[string]int)
	for i := 1; i <= 52; i++ {
		priorities[toCharStr(i)] = i
	}
	return priorities
}

func main() {
	part1()
	part2()
}

func part1() {
	sum, priorities := 0, getPriorities()

	readFile, err := os.Open("2022/day3/input")
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		rucksack := strings.Split(fileScanner.Text(), "")
		length := len(rucksack)
		c1, c2 := rucksack[:length/2], rucksack[length/2:]
		sharedItems := lo.Intersect(c1, c2)
		sum += priorities[sharedItems[0]]
	}
	fmt.Println("Part1:", sum)
}

func part2() {
	sum, priorities := 0, getPriorities()

	readFile, err := os.Open("2022/day3/input")
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		r1 := strings.Split(fileScanner.Text(), "")
		fileScanner.Scan()
		r2 := strings.Split(fileScanner.Text(), "")
		fileScanner.Scan()
		r3 := strings.Split(fileScanner.Text(), "")

		badge := lo.Intersect(r1, lo.Intersect(r2, r3))[0]
		sum += priorities[badge]
	}
	fmt.Println("Part2:", sum)
}
