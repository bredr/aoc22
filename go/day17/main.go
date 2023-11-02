package main

import (
	"fmt"
	"os"
	"strings"
)

type Shape struct {
	Mask   [][]bool
	Height int
	Width  int
}

func main() {
	shape1 := Shape{Mask: [][]bool{{true, true, true, true}}, Height: 1, Width: 4}
	shape2 := Shape{Mask: [][]bool{{false, true, false}, {true, true, true}, {false, true, false}}, Height: 3, Width: 3}
	shape3 := Shape{Mask: [][]bool{{true, true, true}, {false, false, true}, {false, false, true}}, Height: 3, Width: 3}
	shape4 := Shape{Mask: [][]bool{{true}, {true}, {true}, {true}}, Height: 4, Width: 1}
	shape5 := Shape{Mask: [][]bool{{true, true}, {true, true}}, Height: 2, Width: 2}
	shapes := []Shape{shape1, shape2, shape3, shape4, shape5}

	b, err := os.ReadFile("day17/input") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	var jets []int
	for _, char := range strings.Split(str, "") {
		if char == ">" {
			jets = append(jets, 1)
		} else if char == "<" {
			jets = append(jets, -1)
		}
	}

	var chamber [][7]bool

	highest := func() int {
		for idx, row := range chamber {
			contains := false
			for _, point := range row {
				if point {
					contains = true
					break
				}
			}
			if !contains {
				return idx
			}
		}
		return len(chamber)
	}
	print_chamber := func() {
		for i := len(chamber) - 1; i >= 0; i-- {
			row := chamber[i]
			fmt.Print("|")
			for _, point := range row {
				if point {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Print("|\n")
		}
	}
	step := 0
	rocks := 0
all:
	for {
		for _, shape := range shapes {
			// create shape and initial position
			left_x := 2
			h := highest()
			bottom_y := 3 + h
			for {
				jet := jets[step%len(jets)]
				// move by jet
				if left_x+jet >= 0 && left_x+jet+shape.Width <= 7 {
					isBlocked := false
				shape:
					for x := 0; x < shape.Width; x++ {
						for y := 0; y < shape.Height; y++ {
							if bottom_y+y < len(chamber) {
								if shape.Mask[y][x] && chamber[bottom_y+y][left_x+jet+x] {
									isBlocked = true
									break shape
								}
							}
						}
					}
					if !isBlocked {
						left_x += jet
					}
				}

				// move down
				isBlocked := false
				if bottom_y-1 >= 0 {
				shapedown:
					for x := 0; x < shape.Width; x++ {
						for y := 0; y < shape.Height; y++ {
							if bottom_y+y-1 < len(chamber) {
								if shape.Mask[y][x] && chamber[bottom_y+y-1][left_x+x] {
									isBlocked = true
									break shapedown
								}
							}
						}
					}
				} else {
					isBlocked = true
				}
				step++
				if !isBlocked {
					bottom_y -= 1
				} else {
					break
				}
			}
			toAdd := shape.Height + bottom_y - len(chamber)
			for j := 0; j <= toAdd; j++ {
				chamber = append(chamber, [7]bool{})
			}
			for x := 0; x < shape.Width; x++ {
				for y := 0; y < shape.Height; y++ {
					if shape.Mask[y][x] {
						chamber[y+bottom_y][x+left_x] = true
					}
				}
			}
			if rocks < 3 {
				print_chamber()
				fmt.Println()
			}
			rocks++
			if rocks == 2022 {
				break all
			}
		}
	}
	fmt.Println("part1=", highest())
}
