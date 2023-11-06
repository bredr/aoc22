package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type MixedCost struct {
	Ore      int
	Clay     int
	Obsidean int
}

type Blueprint struct {
	ID                int
	OreRobotCost      int
	ClayRobotCost     int
	ObsideanRobotCost MixedCost
	GeodeRobotCost    MixedCost
	MaxOre            int
	MaxClay           int
	MaxObsidian       int
}

type State struct {
	OreRobots      int
	ClayRobots     int
	ObsideanRobots int
	GeodeRobots    int
	Ore            int
	Clay           int
	Obsidean       int
	Geode          int
}

func (b *Blueprint) MaxGeodes(minutes int, state *State) int {
	if minutes == 0 {
		return state.Geode
	}
	clay := state.Clay + state.ClayRobots
	ore := state.Ore + state.OreRobots
	obsidean := state.Obsidean + state.ObsideanRobots
	geode := state.Geode + state.GeodeRobots
	var options []int
	if state.Ore >= b.ClayRobotCost && (state.ClayRobots*minutes+state.Clay < minutes*b.MaxClay) {
		options = append(options, b.MaxGeodes(minutes-1, &State{Clay: clay, Ore: ore - b.ClayRobotCost, Obsidean: obsidean, Geode: geode, OreRobots: state.OreRobots, ClayRobots: state.ClayRobots + 1, ObsideanRobots: state.ObsideanRobots, GeodeRobots: state.GeodeRobots}))
	}
	if state.Ore >= b.OreRobotCost && (state.OreRobots*minutes+state.Ore < minutes*b.MaxOre) {
		options = append(options, b.MaxGeodes(minutes-1, &State{Clay: clay, Ore: ore - b.OreRobotCost, Obsidean: obsidean, Geode: geode, OreRobots: state.OreRobots + 1, ClayRobots: state.ClayRobots, ObsideanRobots: state.ObsideanRobots, GeodeRobots: state.GeodeRobots}))
	}
	if state.Ore >= b.ObsideanRobotCost.Ore && state.Clay >= b.ObsideanRobotCost.Clay && (state.ObsideanRobots*minutes+state.Obsidean < minutes*b.MaxObsidian) {
		options = append(options, b.MaxGeodes(minutes-1, &State{Clay: clay - b.ObsideanRobotCost.Clay, Ore: ore - b.ObsideanRobotCost.Ore, Obsidean: obsidean, Geode: geode, OreRobots: state.OreRobots, ClayRobots: state.ClayRobots, ObsideanRobots: state.ObsideanRobots + 1, GeodeRobots: state.GeodeRobots}))
	}
	if state.Ore >= b.GeodeRobotCost.Ore && state.Obsidean >= b.GeodeRobotCost.Obsidean {
		options = append(options, b.MaxGeodes(minutes-1, &State{Clay: clay, Ore: ore - b.GeodeRobotCost.Ore, Obsidean: obsidean - b.GeodeRobotCost.Obsidean, Geode: geode, OreRobots: state.OreRobots, ClayRobots: state.ClayRobots, ObsideanRobots: state.ObsideanRobots, GeodeRobots: state.GeodeRobots + 1}))
	}
	if state.OreRobots*minutes+state.Ore < b.MaxOre*minutes || len(options) == 0 {
		options = append(options, b.MaxGeodes(minutes-1, &State{Clay: clay, Ore: ore, Obsidean: obsidean, Geode: geode, OreRobots: state.OreRobots, ClayRobots: state.ClayRobots, ObsideanRobots: state.ObsideanRobots, GeodeRobots: state.GeodeRobots}))
	}
	return slices.Max(options)
}

func main() {
	b, err := os.ReadFile("day19/input") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	r := regexp.MustCompile(`\d+`)
	var blueprints []Blueprint
	for _, line := range strings.Split(str, "\n") {
		var numbers []int
		for _, match := range r.FindAllString(line, -1) {
			n, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, n)
		}
		blueprints = append(blueprints, Blueprint{
			ID:                numbers[0],
			OreRobotCost:      numbers[1],
			ClayRobotCost:     numbers[2],
			ObsideanRobotCost: MixedCost{Ore: numbers[3], Clay: numbers[4]},
			GeodeRobotCost:    MixedCost{Ore: numbers[5], Obsidean: numbers[6]},
			MaxOre:            max(numbers[1], numbers[2], numbers[3], numbers[5]),
			MaxClay:           numbers[4],
			MaxObsidian:       numbers[6],
		})
	}
	part1 := 0
	for _, blueprint := range blueprints {
		maxGeodes := blueprint.MaxGeodes(24, &State{OreRobots: 1})
		part1 += blueprint.ID * maxGeodes
	}
	fmt.Println("part1=", part1)
}
