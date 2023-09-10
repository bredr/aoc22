package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	readFile, err := os.Open("day10/input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	x := 1
	cycle := 0
	sum := 0
	cycles := []int{20, 60, 100, 140, 180, 220}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Contains(line, "noop") {
			cycle++
			if slices.Contains(cycles, cycle) {
				sum += cycle * x
			}
		} else {
			addx := strings.Split(line, " ")
			v, _ := strconv.Atoi(addx[1])
			cycle++
			if slices.Contains(cycles, cycle) {
				sum += cycle * x
			}
			cycle++
			if slices.Contains(cycles, cycle) {
				sum += cycle * x
			}
			x += v
		}
	}
	fmt.Printf("part 1: %d\n", sum)
	readFile.Close()
}
