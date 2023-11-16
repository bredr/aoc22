package main

import (
	"fmt"
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
	fmt.Println(knownMonkeys["root"])
}
