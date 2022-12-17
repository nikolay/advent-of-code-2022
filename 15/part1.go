package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)
	var pairs []Pair
	minX, maxX := 0, 0

	scanner := bufio.NewScanner(file)
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
		left, right := sensorX-distance, sensorX+distance
		if len(pairs) == 0 {
			minX, maxX = left, right
		} else {
			if left < minX {
				minX = left
			}
			if right > maxX {
				maxX = right
			}
		}
		pairs = append(pairs, Pair{sensor, beacon, distance})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0
	y := 2000000
	for x := minX; x <= maxX; x++ {
		if !Validate(Coord{x, y}, &pairs) {
			count++
		}
	}
	return count
}
