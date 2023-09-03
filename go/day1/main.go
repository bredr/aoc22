package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("day1/input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	calories := 0
	maxCalories := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			v, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			calories += v
		} else {
			if calories > maxCalories {
				maxCalories = calories
			}
			calories = 0
		}
	}
	fmt.Printf("max calories = %d\n", maxCalories)

	readFile.Close()
}
