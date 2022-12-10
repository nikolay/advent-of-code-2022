package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Screen [6][40]bool

func Draw(screen *Screen, cycle *int, x int) {
	c := *cycle - 1
	row, col := c/40, c%40
	if col >= x-1 && col <= x+1 {
		(*screen)[row][col] = true
	}
	*cycle++
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	x := 1
	cycle := 1
	screen := Screen{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		fields := strings.Fields(line)
		switch fields[0] {
		case "noop":
			Draw(&screen, &cycle, x)
		case "addx":
			Draw(&screen, &cycle, x)
			Draw(&screen, &cycle, x)
			v, _ := strconv.Atoi(fields[1])
			x += v
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for row := 0; row < len(screen); row++ {
		s := ""
		for col := 0; col < len(screen[row]); col++ {
			if screen[row][col] {
				s += "#"
			} else {
				s += "."
			}
		}
		fmt.Println(s)
	}
}
