package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("day6/input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		signal := strings.Split(line, "")
		var last4Buffer []string
		for i, char := range signal {
			if len(last4Buffer) == 4 {
				if allUnique(last4Buffer) {
					fmt.Printf("part1 index %d\n", i)
					break
				}
				last4Buffer = last4Buffer[1:]
			}
			last4Buffer = append(last4Buffer, char)
		}

		var last14Buffer []string
		for i, char := range signal {
			if len(last14Buffer) == 14 {
				if allUnique(last14Buffer) {
					fmt.Printf("part2 index %d\n", i)
					break
				}
				last14Buffer = last14Buffer[1:]
			}
			last14Buffer = append(last14Buffer, char)
		}

	}

	readFile.Close()
}

func allUnique(xx []string) bool {
	for i, x := range xx {
		for j, y := range xx {
			if i != j && x == y {
				return false
			}
		}
	}
	return true
}
