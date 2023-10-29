package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("day16/input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	valves := make(map[string]int)
	flowRateRe := regexp.MustCompile(`\d+`)
	valveRe := regexp.MustCompile(`\s[A-Z][A-Z]\s`)
	paths := make(map[string][]string)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitting := strings.Split(line, ";")
		valve := strings.TrimSpace(valveRe.FindString(splitting[0]))
		flow, _ := strconv.Atoi(flowRateRe.FindString(splitting[0]))
		valves[valve] = flow
		if strings.Contains(splitting[1], "valves") {
			paths[valve] = strings.Split(strings.Split(splitting[1], "valves ")[1], ", ")
		} else {
			paths[valve] = []string{strings.Split(splitting[1], "valve ")[1]}
		}
	}
	readFile.Close()

	// calculate the distance between non broken valves
	shortestPath := func(a, b string) []string {
		var queue []string
		queue = append(queue, a)
		explored := make(map[string]struct{})
		explored[a] = struct{}{}
		parents := make(map[string]string)
		for {
			if len(queue) == 0 {
				break
			}
			v := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			if v == b {
				path := []string{b}
				for {
					if parent, ok := parents[path[len(path)-1]]; ok {
						path = append(path, parent)
					} else {
						return path
					}
				}
			}
			for _, w := range paths[v] {
				if _, ok := explored[w]; !ok {
					explored[w] = struct{}{}
					queue = append(queue, w)
					parents[w] = v
				}
			}

		}
		return make([]string, 31)
	}

	distanceLookups := make(map[string]map[string]int)
	for a := range valves {
		if valves[a] == 0 && a != "AA" {
			continue
		}
		distanceLookups[a] = make(map[string]int)
		for b := range valves {
			if a == b || valves[b] == 0 {
				continue
			}
			path := shortestPath(a, b)
			distanceLookups[a][b] = len(path) - 1
		}
	}

	var f func(currentPosition string, tLeft int, opened map[string]struct{}, released int) int

	f = func(currentPosition string, tLeft int, opened map[string]struct{}, released int) int {
		if tLeft == 0 {
			return released
		}
		var options []int
		for next, tMinus := range distanceLookups[currentPosition] {
			if tLeft-tMinus-1 >= 0 {
				if _, ok := opened[next]; !ok {
					o := copyMap(opened)
					o[next] = struct{}{}
					options = append(options, f(next, tLeft-tMinus-1, o, released+valves[next]*(tLeft-tMinus-1)))
				}
			}
		}
		if len(options) == 0 {
			return released
		}
		return slices.Max(options)
	}

	result := f("AA", 30, make(map[string]struct{}), 0)
	fmt.Println("part1 =", result)

}

func copyMap(m map[string]struct{}) map[string]struct{} {
	out := make(map[string]struct{})
	for k := range m {
		out[k] = struct{}{}
	}
	return out
}
