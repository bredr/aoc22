package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open("day1/input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var calories []int
	subTotal := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			v, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			subTotal += v
		} else {
			calories = append(calories, subTotal)
			subTotal = 0
		}
	}
	sort.Ints(calories)

	fmt.Printf("max calories = %d\n", calories[len(calories)-1])
	fmt.Printf("total top 3 calories = %d\n", calories[len(calories)-1]+calories[len(calories)-2]+calories[len(calories)-3])

	readFile.Close()
}
