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

	rocks := make(map[[2]int]struct{})
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
					rocks[[2]int{x1, y}] = struct{}{}
				}
			} else {
				for x := min(x1, x2); x <= max(x1, x2); x++ {
					rocks[[2]int{x, y1}] = struct{}{}
				}
			}
		}
	}
	grains := make(map[[2]int]struct{})
outerLoop:
	for {
		current := [2]int{500, 0}
		for {
			if current[1] > minY+1 {
				break outerLoop
			}
			// bloked down
			nextPossible := [2]int{current[0], current[1] + 1}
			_, rocksBlocked := rocks[nextPossible]
			_, grainsBlocked := grains[nextPossible]
			if !rocksBlocked && !grainsBlocked {
				current = nextPossible
				continue
			}
			// blocked down-left
			nextPossible = [2]int{current[0] - 1, current[1] + 1}
			_, rocksBlocked = rocks[nextPossible]
			_, grainsBlocked = grains[nextPossible]
			if !rocksBlocked && !grainsBlocked {
				current = nextPossible
				continue
			}
			// blocked down-right
			nextPossible = [2]int{current[0] + 1, current[1] + 1}
			_, rocksBlocked = rocks[nextPossible]
			_, grainsBlocked = grains[nextPossible]
			if !rocksBlocked && !grainsBlocked {
				current = nextPossible
				continue
			}
			/// stop
			break
		}
		grains[current] = struct{}{}
	}
	printLayout(rocks, grains)

	fmt.Println("Total grains", len(grains))

	grains2 := make(map[[2]int]struct{})
	for {
		current := [2]int{500, 0}
		for {
			// bloked down
			nextPossible := [2]int{current[0], current[1] + 1}
			_, rocksBlocked := rocks[nextPossible]
			_, grainsBlocked := grains2[nextPossible]
			bottomBlocked := nextPossible[1] == minY+2
			if !rocksBlocked && !bottomBlocked && !grainsBlocked {
				current = nextPossible
				continue
			}
			// blocked down-left
			nextPossible = [2]int{current[0] - 1, current[1] + 1}
			_, rocksBlocked = rocks[nextPossible]
			_, grainsBlocked = grains2[nextPossible]
			bottomBlocked = nextPossible[1] == minY+2
			if !rocksBlocked && !bottomBlocked && !grainsBlocked {
				current = nextPossible
				continue
			}
			// blocked down-right
			nextPossible = [2]int{current[0] + 1, current[1] + 1}
			_, rocksBlocked = rocks[nextPossible]
			_, grainsBlocked = grains2[nextPossible]
			bottomBlocked = nextPossible[1] == minY+2
			if !rocksBlocked && !bottomBlocked && !grainsBlocked {
				current = nextPossible
				continue
			}
			/// stop
			break
		}
		grains2[current] = struct{}{}
		if current[0] == 500 && current[1] == 0 {
			break
		}
	}
	printLayout(rocks, grains2)

	fmt.Println("Total grains, part 2", len(grains2))
}

func printLayout(rocks map[[2]int]struct{}, grains map[[2]int]struct{}) {
	minX := 500
	maxX := 500
	maxY := 0
	for rock := range rocks {
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
	for grain := range grains {
		x := grain[0]
		y := grain[1]
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
	for y := 0; y <= maxY+2; y++ {
	coord:
		for x := minX; x <= maxX; x++ {
			if _, ok := rocks[[2]int{x, y}]; ok {
				fmt.Print("#")
				continue coord
			}
			if _, ok := grains[[2]int{x, y}]; ok {
				fmt.Print("+")
				continue coord
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
	fmt.Println()

}
