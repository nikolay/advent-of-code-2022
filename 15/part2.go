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

func Diamond(center Coord, size int) []Coord {
	dirs := []Coord{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}}
	i, result := 0, make([]Coord, size*len(dirs))
	pixel := Coord{center.x, center.y - size}
	for _, dir := range dirs {
		for n := size; n > 0; n-- {
			result[i] = pixel
			i++
			pixel = pixel.Add(dir)
		}
	}
	return result
}

func Part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)
	var pairs []Pair
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		matches := r.FindStringSubmatch(line)
		if len(matches) == 0 {
			log.Fatalf("invalid command: %v", line)
		}
		sensorX, _ := strconv.Atoi(matches[1])
		sensorY, _ := strconv.Atoi(matches[2])
		beaconX, _ := strconv.Atoi(matches[3])
		beaconY, _ := strconv.Atoi(matches[4])
		sensor := Coord{sensorX, sensorY}
		beacon := Coord{beaconX, beaconY}
		distance := sensor.Distance(beacon)
		pairs = append(pairs, Pair{sensor, beacon, distance})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	distance := 0
outer:
	for {
		distance++
		for _, pair := range pairs {
			for _, candidate := range Diamond(pair.sensor, pair.distance+distance) {
				if candidate.x >= 0 && candidate.x <= 4000000 && candidate.y >= 0 && candidate.y <= 4000000 {
					if Validate(candidate, &pairs) {
						fmt.Println(candidate.x*4000000 + candidate.y)
						break outer
					}
				}
			}
		}
	}
}
