package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("day2/input")

	if err != nil {
		fmt.Println(err)
	}
	lineToScorePart1 := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	lineToScorePart2 := map[string]int{
		"A X": 0 + 3,
		"A Y": 3 + 1,
		"A Z": 6 + 2,
		"B X": 0 + 1,
		"B Y": 3 + 2,
		"B Z": 6 + 3,
		"C X": 0 + 2,
		"C Y": 3 + 3,
		"C Z": 6 + 1,
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	scorePart1 := 0
	scorePart2 := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		scorePart1 += lineToScorePart1[line]
		scorePart2 += lineToScorePart2[line]
	}
	fmt.Printf("part1 score = %d\n", scorePart1)
	fmt.Printf("part2 score = %d\n", scorePart2)

}
