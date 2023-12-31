package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("day4/input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	completeOverlaps := 0
	partialOverlaps := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		pair := strings.Split(line, ",")
		elf1Raw := strings.Split(pair[0], "-")
		elf1 := make([]int, 2)
		for i, x := range elf1Raw {
			elf1[i], _ = strconv.Atoi(x)
		}
		elf2Raw := strings.Split(pair[1], "-")
		elf2 := make([]int, 2)
		for i, x := range elf2Raw {
			elf2[i], _ = strconv.Atoi(x)
		}
		if (elf1[0] >= elf2[0] && elf1[1] <= elf2[1]) || (elf2[0] >= elf1[0] && elf2[1] <= elf1[1]) {
			completeOverlaps += 1
		}
		if (elf1[0] >= elf2[0] && elf1[0] <= elf2[1]) || (elf2[0] >= elf1[0] && elf2[0] <= elf1[1]) {
			partialOverlaps += 1
		}
	}
	log.Printf("Part 1 overlaps = %d\n", completeOverlaps)
	log.Printf("Part 2 all overlaps = %d\n", partialOverlaps)

}
