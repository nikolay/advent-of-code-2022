package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Check(cycle *int, x int, sum *int) {
	if (*cycle-20)%40 == 0 {
		*sum += *cycle * x
	}
	*cycle++
}

func Part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x := 1
	cycle := 1
	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		fields := strings.Fields(line)
		switch fields[0] {
		case "noop":
			Check(&cycle, x, &sum)
		case "addx":
			Check(&cycle, x, &sum)
			Check(&cycle, x, &sum)
			v, _ := strconv.Atoi(fields[1])
			x += v
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}
