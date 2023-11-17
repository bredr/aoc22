package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Elf struct {
	X        int
	Y        int
	Proposal *Elf
}

func main() {

	b, err := os.ReadFile("day23/input") // just pass the file name
	if err != nil {
		log.Fatal(err)
	}

	str := string(b) // convert content to a 'string'
	raw := strings.Split(str, "\n")

	var elves []*Elf
	for i, row := range raw {
		y := len(raw) - i
		for x, char := range strings.Split(row, "") {
			if char == "#" {
				elves = append(elves, &Elf{X: x, Y: y})
			}
		}
	}

	for i := 0; i < 10; i++ {
		makeProposals(elves, i%4)
		removeDuplicateProposals(elves)
		moveElves(elves)
	}
	minX := 1000
	maxX := 0
	minY := 1000
	maxY := 0
	for _, elf := range elves {
		if elf.X > maxX {
			maxX = elf.X
		}
		if elf.X < minX {
			minX = elf.X
		}
		if elf.Y > maxY {
			maxY = elf.Y
		}
		if elf.Y < minY {
			minY = elf.Y
		}
	}
	fmt.Println("part1 = ", (maxX-minX+1)*(maxY-minY+1)-len(elves))

}

func makeProposals(elves []*Elf, direction int) {
	for i, elf := range elves {
		x := elf.X
		y := elf.Y
		proposeNorth := true
		proposeSouth := true
		proposeEast := true
		proposeWest := true
		surrounded := 0
		for j, other := range elves {
			if !proposeNorth && !proposeSouth && !proposeEast && !proposeWest {
				break
			}
			N := (other.X == x && other.Y == y+1)
			NE := (other.X == x+1 && other.Y == y+1)
			NW := (other.X == x-1 && other.Y == y+1)
			E := (other.X == x+1 && other.Y == y)
			SE := (other.X == x+1 && other.Y == y-1)
			S := (other.X == x && other.Y == y-1)
			SW := (other.X == x-1 && other.Y == y-1)
			W := (other.X == x-1 && other.Y == y)
			if i == j {
				continue
			}
			if proposeNorth && (N || NW || NE) {
				proposeNorth = false
			}
			if proposeSouth && (S || SW || SE) {
				proposeSouth = false
			}
			if proposeEast && (E || SE || NE) {
				proposeEast = false
			}
			if proposeWest && (W || SW || NW) {
				proposeWest = false
			}
			if N || S || E || W || NW || NE || SE || SW {
				surrounded++
			}
		}
		if surrounded == 0 {
			continue
		}
		switch direction {
		case 0:
			if proposeNorth {
				elf.Proposal = &Elf{X: x, Y: y + 1}
				continue
			}
			if proposeSouth {
				elf.Proposal = &Elf{X: x, Y: y - 1}
				continue
			}
			if proposeWest {
				elf.Proposal = &Elf{X: x - 1, Y: y}
				continue
			}
			if proposeEast {
				elf.Proposal = &Elf{X: x + 1, Y: y}
				continue
			}
		case 1:
			if proposeSouth {
				elf.Proposal = &Elf{X: x, Y: y - 1}
				continue
			}
			if proposeWest {
				elf.Proposal = &Elf{X: x - 1, Y: y}
				continue
			}
			if proposeEast {
				elf.Proposal = &Elf{X: x + 1, Y: y}
				continue
			}
			if proposeNorth {
				elf.Proposal = &Elf{X: x, Y: y + 1}
				continue
			}
		case 2:
			if proposeWest {
				elf.Proposal = &Elf{X: x - 1, Y: y}
				continue
			}
			if proposeEast {
				elf.Proposal = &Elf{X: x + 1, Y: y}
				continue
			}
			if proposeNorth {
				elf.Proposal = &Elf{X: x, Y: y + 1}
				continue
			}
			if proposeSouth {
				elf.Proposal = &Elf{X: x, Y: y - 1}
				continue
			}
		case 3:
			if proposeEast {
				elf.Proposal = &Elf{X: x + 1, Y: y}
				continue
			}
			if proposeNorth {
				elf.Proposal = &Elf{X: x, Y: y + 1}
				continue
			}
			if proposeSouth {
				elf.Proposal = &Elf{X: x, Y: y - 1}
				continue
			}
			if proposeWest {
				elf.Proposal = &Elf{X: x - 1, Y: y}
				continue
			}
		}
		elf.Proposal = nil
	}
}

func removeDuplicateProposals(elves []*Elf) {
	for i, elf := range elves {
		if elf.Proposal != nil {
			x := elf.Proposal.X
			y := elf.Proposal.Y
			for j, other := range elves {
				if i == j {
					continue
				}
				if other.Proposal != nil {
					if other.Proposal.X == x && other.Proposal.Y == y {
						other.Proposal = nil
						elf.Proposal = nil
					}
				}
			}
		}
	}
}

func moveElves(elves []*Elf) {
	for _, elf := range elves {
		if elf.Proposal != nil {
			elf.X = elf.Proposal.X
			elf.Y = elf.Proposal.Y
			elf.Proposal = nil
		}
	}
}

func printElves(elves []*Elf) {
	minX := 1000
	maxX := 0
	minY := 1000
	maxY := 0
	for _, elf := range elves {
		if elf.X > maxX {
			maxX = elf.X
		}
		if elf.X < minX {
			minX = elf.X
		}
		if elf.Y > maxY {
			maxY = elf.Y
		}
		if elf.Y < minY {
			minY = elf.Y
		}
	}

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			empty := true
			for _, elf := range elves {
				if elf.X == x && elf.Y == y {
					fmt.Print("#")
					empty = false
					break
				}
			}
			if empty {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
