package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type ToSolve struct {
	MonkeyLeft  string
	MonkeyRight string
	Operation   string
}

func main() {
	b, err := os.ReadFile("day21/input") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	raw := strings.Split(str, "\n")
	knownMonkeys := make(map[string]int)
	monkeysToSolve := make(map[string]ToSolve)
	for _, line := range raw {
		splitLine := strings.Split(line, ":")
		monkey := splitLine[0]
		value, err := strconv.Atoi(strings.TrimSpace(splitLine[1]))
		if err != nil {
			splitToSolve := strings.Split(strings.TrimSpace(splitLine[1]), " ")
			monkeysToSolve[monkey] = ToSolve{MonkeyLeft: splitToSolve[0], MonkeyRight: splitToSolve[2], Operation: splitToSolve[1]}
		} else {
			knownMonkeys[monkey] = value
		}
	}
	for {
		for monkey, toSolve := range monkeysToSolve {
			valueLeft, ok := knownMonkeys[toSolve.MonkeyLeft]
			if !ok {
				continue
			}
			valueRight, ok := knownMonkeys[toSolve.MonkeyRight]
			if !ok {
				continue
			}
			switch toSolve.Operation {
			case "*":
				knownMonkeys[monkey] = valueLeft * valueRight
			case "-":
				knownMonkeys[monkey] = valueLeft - valueRight
			case "/":
				knownMonkeys[monkey] = valueLeft / valueRight
			case "+":
				knownMonkeys[monkey] = valueLeft + valueRight
			}
			delete(monkeysToSolve, monkey)
		}
		if len(monkeysToSolve) == 0 {
			break
		}
	}
	fmt.Println("part1 = ", knownMonkeys["root"])

	funcPart2 := func(humn int) int {
		knownMonkeys := make(map[string]int)
		monkeysToSolve := make(map[string]ToSolve)
		var monkeyRootA, monkeyRootB string
		for _, line := range raw {
			splitLine := strings.Split(line, ":")
			monkey := splitLine[0]
			value, err := strconv.Atoi(strings.TrimSpace(splitLine[1]))
			if err != nil {
				splitToSolve := strings.Split(strings.TrimSpace(splitLine[1]), " ")
				if monkey == "root" {
					monkeyRootA = splitToSolve[0]
					monkeyRootB = splitToSolve[2]
				} else {
					monkeysToSolve[monkey] = ToSolve{MonkeyLeft: splitToSolve[0], MonkeyRight: splitToSolve[2], Operation: splitToSolve[1]}
				}
			} else {
				if monkey == "humn" {
					knownMonkeys[monkey] = humn
				} else {
					knownMonkeys[monkey] = value
				}
			}
		}
		for {
			for monkey, toSolve := range monkeysToSolve {
				valueLeft, ok := knownMonkeys[toSolve.MonkeyLeft]
				if !ok {
					continue
				}
				valueRight, ok := knownMonkeys[toSolve.MonkeyRight]
				if !ok {
					continue
				}
				switch toSolve.Operation {
				case "*":
					knownMonkeys[monkey] = valueLeft * valueRight
				case "-":
					knownMonkeys[monkey] = valueLeft - valueRight
				case "/":
					knownMonkeys[monkey] = valueLeft / valueRight
				case "+":
					knownMonkeys[monkey] = valueLeft + valueRight
				}
				delete(monkeysToSolve, monkey)
			}
			if len(monkeysToSolve) == 0 {
				return knownMonkeys[monkeyRootA] - knownMonkeys[monkeyRootB]
			}
		}
	}

	humnOld := 0
	humn := 100
	gamma := 0.01
	for {
		if math.Abs(float64(humnOld-humn)) < 1 {
			fmt.Println("part2 = ", humn)
			break
		}
		diff := funcPart2(humn)
		humnOld = humn
		humn += int(gamma * float64(diff))
	}

}
