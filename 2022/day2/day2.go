package main

import (
	"bufio"
	"fmt"
	"os"
)

var test = "2022/day2/test"
var input = "2022/day2/input"

func main() {
	rpsRounds := buildInput()
	p1Score := 0
	p2Score := 0
	for _, v := range rpsRounds {
		elf := v[:1]
		you := v[2:]
		p1Score += Fight(elf, you)
		p2Score += UltraFight(elf, you)
	}
	fmt.Println("p1 Score:", p1Score)
	fmt.Println("p2 Score:", p2Score)
}

func buildInput() []string {
	readFile, err := os.Open(input)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	rpsRounds := make([]string, 0)

	for fileScanner.Scan() {
		var strVar = fileScanner.Text()
		rpsRounds = append(rpsRounds, strVar)
	}
	readFile.Close()
	return rpsRounds
}
func Fight(elf string, you string) int {
	if elf == "A" {
		if you == "X" {
			return 1 + 3
		} else if you == "Y" {
			return 2 + 6
		} else {
			return 3 + 0
		}
	} else if elf == "B" {
		if you == "X" {
			return 1 + 0
		} else if you == "Y" {
			return 2 + 3
		} else {
			return 3 + 6
		}
	} else if elf == "C" {
		if you == "X" {
			return 1 + 6
		} else if you == "Y" {
			return 2 + 0
		} else {
			return 3 + 3
		}
	}
	return 0
}
func UltraFight(elf string, you string) int { // X = Lose, Y = Draw and Z = Win
	if elf == "A" { // Rock
		if you == "X" {
			return 3 + 0
		} else if you == "Y" {
			return 1 + 3
		} else {
			return 2 + 6
		}
	} else if elf == "B" { // Paper
		if you == "X" {
			return 1 + 0
		} else if you == "Y" {
			return 2 + 3
		} else {
			return 3 + 6
		}
	} else if elf == "C" { // Scissors
		if you == "X" {
			return 2 + 0
		} else if you == "Y" {
			return 3 + 3
		} else {
			return 1 + 6
		}
	}
	return 0
}
