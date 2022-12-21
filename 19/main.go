package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Resources struct {
	ore, clay, obsidian, geode int
}

type Bot struct {
	collects Resources
	costs    Resources
}

type Blueprint struct {
	id   int
	bots []Bot
}

type States []State

type State struct {
	bots      Resources
	inventory Resources
	collected Resources
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (resources *Resources) Add(add Resources, factor int) Resources {
	return Resources{
		resources.ore + add.ore*factor,
		resources.clay + add.clay*factor,
		resources.obsidian + add.obsidian*factor,
		resources.geode + add.geode*factor,
	}
}

func (resources *Resources) Resource(resourceName string) *int {
	switch resourceName {
	case "ore":
		return &resources.ore
	case "clay":
		return &resources.clay
	case "obsidian":
		return &resources.obsidian
	case "geode":
		return &resources.geode
	}
	return nil
}

func (resources *Resources) CanAfford(costs Resources) bool {
	return resources.ore >= costs.ore &&
		resources.clay >= costs.clay &&
		resources.obsidian >= costs.obsidian &&
		resources.geode >= costs.geode
}

func (state State) CompareFitness(champion State) bool {
	weights := Resources{
		1,
		16,
		256,
		1024,
	}
	return (state.collected.ore-champion.collected.ore)*weights.ore+
		(state.collected.clay-champion.collected.clay)*weights.clay+
		(state.collected.obsidian-champion.collected.obsidian)*weights.obsidian+
		(state.collected.geode-champion.collected.geode)*weights.geode < 0
}

func (states States) Len() int {
	return len(states)
}

func (states States) Less(i, j int) bool {
	return states[i].CompareFitness(states[j])
}

func (states States) Swap(i, j int) {
	states[i], states[j] = states[j], states[i]
}

func Solve(blueprint Blueprint, time int, maxStates int) int {
	states := States{{Resources{ore: 1}, Resources{}, Resources{}}}
	for minute := 0; minute < time; minute++ {
		var newStates States
		for _, state := range states {
			inventory := state.inventory.Add(state.bots, 1)
			collected := state.collected.Add(state.bots, 1)
			newStates = append(newStates, State{
				state.bots,
				inventory,
				collected,
			})

			for _, bot := range blueprint.bots {
				if state.inventory.CanAfford(bot.costs) {
					newStates = append(newStates, State{
						state.bots.Add(bot.collects, 1),
						inventory.Add(bot.costs, -1),
						collected,
					})
				}
			}
		}

		sort.Sort(sort.Reverse(newStates))
		states = newStates[:Min(maxStates, len(newStates))]
	}
	maxGeodes := 0
	for _, state := range states {
		maxGeodes = Max(state.collected.geode, maxGeodes)
	}
	return maxGeodes
}

func Solve1(blueprints []Blueprint) int {
	result := 0
	for _, blueprint := range blueprints {
		result += blueprint.id * Solve(blueprint, 24, 1024)
	}
	return result
}

func Solve2(blueprints []Blueprint) int {
	result := 1
	for _, blueprint := range blueprints {
		result *= Solve(blueprint, 32, 16384)
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r1 := regexp.MustCompile(`^Each ([a-z]+) robot costs (\d+) ([a-z]+)$`)
	r2 := regexp.MustCompile(`^Each ([a-z]+) robot costs (\d+) ([a-z]+) and (\d+) ([a-z]+)$`)

	var blueprints []Blueprint

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		colon := strings.Index(line, ":")
		fields := strings.Fields(line[:colon])
		id, _ := strconv.Atoi(strings.TrimSpace(fields[1]))
		blueprint := Blueprint{id, []Bot{}}

		fields = strings.Split(strings.TrimSpace(line[colon+1:]), ".")
		for _, field := range fields {
			f := strings.TrimSpace(field)
			if len(f) > 0 {
				bot := Bot{Resources{}, Resources{}}

				matches := r2.FindStringSubmatch(f)
				if len(matches) == 0 {
					matches = r1.FindStringSubmatch(f)
					if len(matches) == 0 {
						log.Fatalf("invalid comamnd: %v", f)
					}
				} else {
					cost, _ := strconv.Atoi(matches[4])
					*bot.costs.Resource(matches[5]) += cost
				}
				cost, _ := strconv.Atoi(matches[2])
				*bot.costs.Resource(matches[3]) += cost

				*bot.collects.Resource(matches[1])++
				blueprint.bots = append(blueprint.bots, bot)
			}
		}

		blueprints = append(blueprints, blueprint)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(Solve1(blueprints))
	fmt.Println(Solve2(blueprints[:Min(3, len(blueprints))]))
}
