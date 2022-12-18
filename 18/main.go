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

func (coord Coord) Move(dx, dy, dz int) Coord {
	return Coord{coord.x + dx, coord.y + dy, coord.z + dz}
}

func (coord Coord) Inside(min, max Coord) bool {
	return coord.x >= min.x && coord.x <= max.x &&
		coord.y >= min.y && coord.y <= max.y &&
		coord.z >= min.z && coord.z <= max.z
}

func (coord Coord) Adjacent() []Coord {
	return []Coord{
		coord.Move(-1, 0, 0),
		coord.Move(+1, 0, 0),
		coord.Move(0, -1, 0),
		coord.Move(0, +1, 0),
		coord.Move(0, 0, -1),
		coord.Move(0, 0, +1),
	}
}

func (coord Coord) Min(min Coord) Coord {
	return Coord{Min(coord.x, min.x), Min(coord.y, min.y), Min(coord.z, min.z)}
}

func (coord Coord) Max(min Coord) Coord {
	return Coord{Max(coord.x, min.x), Max(coord.y, min.y), Max(coord.z, min.z)}
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

func Solve1(cubes []Coord) int {
	space := map[Coord]bool{}
	for _, coord := range cubes {
		space[coord] = true
	}
	surface := 0
	for _, coord := range cubes {
		sides := 0
		for _, adjacent := range coord.Adjacent() {
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
		min = coord.Min(min)
		max = coord.Max(max)
	}
	min = min.Move(-1, -1, -1)
	max = max.Move(+1, +1, +1)

	step := 1
	water := map[Coord]int{Coord{min.x, min.y, min.z}: step}
	for {
		found := false
		batch := map[Coord]int{}
		for coord, v := range water {
			if v == step {
				for _, adjacent := range coord.Adjacent() {
					if adjacent.Inside(min, max) {
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
		for _, adjacent := range coord.Adjacent() {
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
