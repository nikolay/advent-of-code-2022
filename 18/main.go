package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y, z int
}

func Adjacent(coord Coord) []Coord {
	return []Coord{
		{coord.x - 1, coord.y, coord.z},
		{coord.x + 1, coord.y, coord.z},
		{coord.x, coord.y - 1, coord.z},
		{coord.x, coord.y + 1, coord.z},
		{coord.x, coord.y, coord.z - 1},
		{coord.x, coord.y, coord.z + 1},
	}
}

func Solve1(cubes []Coord) int {
	space := map[Coord]bool{}
	for _, coord := range cubes {
		space[coord] = true
	}
	surface := 0
	for _, coord := range cubes {
		sides := 0
		for _, adjacent := range Adjacent(coord) {
			if _, ok := space[adjacent]; !ok {
				sides++
			}
		}
		surface += sides
	}
	return surface
}

func Solve2(cubes []Coord) int {
	space := map[Coord]bool{}
	min := cubes[0]
	max := min
	for _, coord := range cubes {
		space[coord] = true

		if coord.x < min.x {
			min.x = coord.x
		}
		if coord.y < min.y {
			min.y = coord.y
		}
		if coord.z < min.z {
			min.z = coord.z
		}

		if coord.x > max.x {
			max.x = coord.x
		}
		if coord.y > max.y {
			max.y = coord.y
		}
		if coord.z > max.z {
			max.z = coord.z
		}
	}

	step := 1
	water := map[Coord]int{Coord{min.x - 1, min.y - 1, min.z - 1}: step}
	for {
		found := false
		batch := map[Coord]int{}
		for coord, v := range water {
			if v == step {
				for _, adjacent := range Adjacent(coord) {
					if adjacent.x >= min.x-1 && adjacent.x <= max.x+1 &&
						adjacent.y >= min.y-1 && adjacent.y <= max.y+1 &&
						adjacent.z >= min.z-1 && adjacent.z <= max.z+1 {
						if _, ok := space[adjacent]; ok {
							continue
						}
						if _, ok := batch[adjacent]; ok {
							continue
						}
						if _, ok := water[adjacent]; !ok {
							batch[adjacent] = step + 1
							found = true
						}
					}
				}
			}
		}
		if !found {
			break
		}
		for coord, v := range batch {
			water[coord] = v
		}
		step++
	}

	surface := 0
	for _, coord := range cubes {
		sides := 0
		for _, adjacent := range Adjacent(coord) {
			if _, ok := water[adjacent]; ok {
				sides++
			}
		}
		surface += sides
	}
	return surface
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var cubes []Coord

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		fields := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(fields[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(fields[1]))
		z, _ := strconv.Atoi(strings.TrimSpace(fields[2]))
		cubes = append(cubes, Coord{x, y, z})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(Solve1(cubes))
	fmt.Println(Solve2(cubes))
}
