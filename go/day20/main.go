package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Index int
}

func main() {
	b, err := os.ReadFile("day20/input") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	raw := strings.Split(str, "\n")

	nPositions := len(raw)
	values := make([]int, nPositions)
	for i := range values {
		n, err := strconv.Atoi(raw[i])
		if err != nil {
			panic(err)
		}
		values[i] = n
	}
	part1 := shuffled(values, 1, 1)

	indexOfZero := 0
	for i, v := range part1 {
		if v == 0 {
			indexOfZero = i
			break
		}
	}
	part1Answer := part1[(indexOfZero+1000)%len(part1)] + part1[(indexOfZero+2000)%len(part1)] + part1[(indexOfZero+3000)%len(part1)]
	fmt.Println("part1=", part1Answer)

	part2 := shuffled(values, 811589153, 10)

	indexOfZero = 0
	for i, v := range part2 {
		if v == 0 {
			indexOfZero = i
			break
		}
	}
	part2Answer := part2[(indexOfZero+1000)%len(part2)] + part2[(indexOfZero+2000)%len(part2)] + part2[(indexOfZero+3000)%len(part2)]
	fmt.Println("part2=", part2Answer)

}

func shuffled(values []int, decriptionKey int, rounds int) []int {
	// copy to outputs
	output := make([]Node, len(values))
	for i, v := range values {
		output[i] = Node{Value: v * decriptionKey, Index: i}
	}
	out := make([]int, len(values))
	for round := 0; round < rounds; round++ {
		for i := range values {
			nodeOfInstruction := Node{}
			indexOfInstruction := 0
			for j, node := range output {
				if node.Index == i {
					nodeOfInstruction = node
					indexOfInstruction = j
					break
				}
			}
			instruction := (nodeOfInstruction.Value % (len(values) - 1))
			newIndex := indexOfInstruction + instruction
			if newIndex >= len(values) {
				newIndex = newIndex - len(values) + 1
			} else if newIndex <= 0 {
				newIndex = newIndex + len(values) - 1
			}
			output = append(output[:indexOfInstruction], output[indexOfInstruction+1:]...)
			output = append(output[:newIndex], append([]Node{nodeOfInstruction}, output[newIndex:]...)...)
		}
	}
	for i, v := range output {
		out[i] = v.Value
	}
	return out
}
