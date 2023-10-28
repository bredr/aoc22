package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Pair struct {
	Sensor [2]int
	Beacon [2]int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func hamiltonDistance(a, b [2]int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func main() {

	readFile, err := os.Open("day15/input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	findNumbers := regexp.MustCompile(`-*\d+`)
	var closest []Pair
	for fileScanner.Scan() {
		line := fileScanner.Text()
		matches := findNumbers.FindAllString(line, 4)
		sensorX, _ := strconv.Atoi(matches[0])
		sensorY, _ := strconv.Atoi(matches[1])
		beaconX, _ := strconv.Atoi(matches[2])
		beaconY, _ := strconv.Atoi(matches[3])
		closest = append(closest, Pair{[2]int{sensorX, sensorY}, [2]int{beaconX, beaconY}})
	}
	readFile.Close()

	noBeacon := make(map[[2]int]struct{})
	hasBeacon := make(map[[2]int]struct{})
	yOfInterest := 2000000
	for _, pair := range closest {
		if pair.Sensor[1] == yOfInterest {
			noBeacon[pair.Sensor] = struct{}{}
		}
		if pair.Beacon[1] == yOfInterest {
			hasBeacon[pair.Beacon] = struct{}{}
		}
	}

	for _, pair := range closest {
		maxDistance := hamiltonDistance(pair.Sensor, pair.Beacon)
		if pair.Sensor[1]-maxDistance <= yOfInterest && pair.Sensor[1]+maxDistance >= yOfInterest {
			for x := pair.Sensor[0] - (abs(maxDistance) - abs(yOfInterest-pair.Sensor[1])); x <= pair.Sensor[0]+(abs(maxDistance)-abs(yOfInterest-pair.Sensor[1])); x++ {
				if _, ok := hasBeacon[[2]int{x, yOfInterest}]; !ok {
					noBeacon[[2]int{x, yOfInterest}] = struct{}{}
				}
			}
		}
	}

	fmt.Println("part1=", len(noBeacon))

}
