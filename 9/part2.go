package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	const TailSize = 10

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ropes := [TailSize]Coord{}
	positions := map[string]bool{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		fields := strings.Fields(line)
		delta := ParseDirection(fields[0])
		for steps, _ := strconv.Atoi(fields[1]); steps > 0; steps-- {
			ropes[0] = ropes[0].Add(delta)
			for r := 1; r < TailSize; r++ {
				dist := ropes[r-1].Subtract(ropes[r])
				if Abs(dist.x) > 1 || Abs(dist.y) > 1 {
					ropes[r] = ropes[r].Move(Sign(dist.x), Sign(dist.y))
					if r == TailSize-1 {
						positions[fmt.Sprintf("%v,%v", ropes[r].x, ropes[r].y)] = true
					}
				} else {
					break
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	count := 0
	for range positions {
		count++
	}
	fmt.Println(count)
}
