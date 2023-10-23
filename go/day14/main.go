package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("day14/input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var paths [][][]int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		coords := strings.Split(line, " -> ")
		var path [][]int
		for _, coord := range coords {
			coord := strings.Split(coord, ",")
			x, _ := strconv.Atoi(coord[0])
			y, _ := strconv.Atoi(coord[1])
			path = append(path, []int{x, y})
		}
		paths = append(paths, path)
	}
	readFile.Close()

	var rocks [][]int
	minY := 0
	for _, path := range paths {
		for ix, coord := range path {
			if ix+1 >= len(path) {
				break
			}
			x1 := coord[0]
			y1 := coord[1]
			x2 := path[ix+1][0]
			y2 := path[ix+1][1]
			if y1 > minY {
				minY = y1
			}
			if x1 == x2 {
				for y := min(y1, y2); y <= max(y1, y2); y++ {
					rocks = append(rocks, []int{x1, y})
				}
			} else {
				for x := min(x1, x2); x <= max(x1, x2); x++ {
					rocks = append(rocks, []int{x, y1})
				}
			}
			rocks = append(rocks, coord)
		}
	}
	var grains [][]int
	printLayout(rocks, grains)
outerLoop:
	for {
		current := []int{500, 0}
		for {
			if current[1] > minY+1 {
				break outerLoop
			}
			// bloked down
			nextPossible := []int{current[0], current[1] + 1}
			blocked := false
			for _, rock := range rocks {
				if nextPossible[0] == rock[0] && nextPossible[1] == rock[1] {
					blocked = true
					break
				}
			}
			for _, grain := range grains {
				if nextPossible[0] == grain[0] && nextPossible[1] == grain[1] {
					blocked = true
					break
				}
			}
			if !blocked {
				current = nextPossible
				continue
			}
			// blocked down-left
			blocked = false
			nextPossible = []int{current[0] - 1, current[1] + 1}
			for _, rock := range rocks {
				if nextPossible[0] == rock[0] && nextPossible[1] == rock[1] {
					blocked = true
					break
				}
			}
			for _, grain := range grains {
				if nextPossible[0] == grain[0] && nextPossible[1] == grain[1] {
					blocked = true
					break
				}
			}
			if !blocked {
				current = nextPossible
				continue
			}
			// blocked down-right
			blocked = false
			nextPossible = []int{current[0] + 1, current[1] + 1}
			for _, rock := range rocks {
				if nextPossible[0] == rock[0] && nextPossible[1] == rock[1] {
					blocked = true
					break
				}
			}
			for _, grain := range grains {
				if nextPossible[0] == grain[0] && nextPossible[1] == grain[1] {
					blocked = true
					break
				}
			}
			if !blocked {
				current = nextPossible
				continue
			}
			/// stop
			break
		}
		grains = append(grains, current)
	}
	printLayout(rocks, grains)

	fmt.Println("Total grains", len(grains))
}

func printLayout(rocks [][]int, grains [][]int) {
	minX := 500
	maxX := 500
	maxY := 0
	for _, rock := range rocks {
		x := rock[0]
		y := rock[1]
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x

		}
		if y > maxY {
			maxY = y
		}
	}
	for y := 0; y <= maxY; y++ {
	coord:
		for x := minX; x <= maxX; x++ {
			for _, rock := range rocks {
				if rock[0] == x && rock[1] == y {
					fmt.Print("#")
					continue coord
				}
			}
			for _, grain := range grains {
				if grain[0] == x && grain[1] == y {
					fmt.Print("+")
					continue coord
				}
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
	fmt.Println()

}
