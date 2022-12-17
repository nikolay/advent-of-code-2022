package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Valve struct {
	index    int
	flowRate int
	tunnels  []int
}

type Input struct {
	names   []string
	indices map[string]int
	valves  map[int]Valve
}

func (input *Input) FindValve(name string) int {
	if n, ok := input.indices[name]; ok {
		return n
	}
	n := len(input.names)
	input.names = append(input.names, name)
	input.indices[name] = n
	return n
}

func (input *Input) AddValve(valve Valve) {
	input.valves[valve.index] = valve
}

func GetInput(filename string) (result Input) {
	result.indices = map[string]int{}
	result.valves = map[int]Valve{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(`^Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? (.+)+$`)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		matches := r.FindStringSubmatch(line)
		if len(matches) == 0 {
			log.Fatalf("invalid command: %v", line)
		}
		var valve Valve
		valve.index = result.FindValve(matches[1])
		valve.flowRate, _ = strconv.Atoi(matches[2])
		for _, tunnel := range strings.Split(matches[3], ", ") {
			valve.tunnels = append(valve.tunnels, result.FindValve(tunnel))
		}
		result.AddValve(valve)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	input := GetInput("input.txt")
	fmt.Println(Part1(&input, 30))
	fmt.Println(Part2(&input, 26))
}
