package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	head := Coord{0, 0}
	tail := head
	positions := map[string]bool{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		fields := strings.Fields(line)
		delta := ParseDirection(fields[0])
		for steps, _ := strconv.Atoi(fields[1]); steps > 0; steps-- {
			head = head.Add(delta)
			dist := head.Subtract(tail)
			if Abs(dist.x) > 1 || Abs(dist.y) > 1 {
				tail = tail.Move(Sign(dist.x), Sign(dist.y))
				positions[fmt.Sprintf("%v,%v", tail.x, tail.y)] = true
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	count := 0
	for _ = range positions {
		count++
	}
	fmt.Println(count)
}
