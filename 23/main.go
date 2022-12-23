package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
	x, y int
}

type Grove struct {
	elves []Coord
}

var Zero = Coord{0, 0}

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

func (coord Coord) Move(dx, dy int) Coord {
	return Coord{coord.x + dx, coord.y + dy}
}

func (coord Coord) Add(add Coord) Coord {
	return coord.Move(add.x, add.y)
}

func (grove *Grove) Bounds() (Coord, Coord) {
	var ne, sw Coord
	for _, elf := range grove.elves {
		ne = Coord{Min(ne.x, elf.x), Min(ne.y, elf.y)}
		sw = Coord{Max(sw.x, elf.x), Max(sw.y, elf.y)}
	}
	return ne, sw
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grove Grove

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				grove.elves = append(grove.elves, Coord{x, y})
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	firstDirTry := 0
	round := 1
	for {
		bits := map[Coord]bool{}
		for _, elf := range grove.elves {
			bits[elf] = true
		}
		free := func(coord Coord) bool {
			if v, ok := bits[coord]; ok && v {
				return false
			}
			return true
		}

		var few []int
		for elfNo, elf := range grove.elves {
			count := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if free(elf.Move(dx, dy)) {
						count++
					}
				}
			}
			if count < 8 {
				few = append(few, elfNo)
			}
		}
		if len(few) == 0 {
			// Part 2
			fmt.Println(round)
			break
		}

		destinations := map[Coord][]int{}
		for _, elfNo := range few {
			elf := grove.elves[elfNo]

			var delta Coord
		trying:
			for try := 0; try < 4; try++ {
				switch (firstDirTry + try) % 4 {
				case 0:
					if free(elf.Move(-1, -1)) && free(elf.Move(0, -1)) && free(elf.Move(1, -1)) {
						delta = Coord{0, -1}
						break trying
					}
				case 1:
					if free(elf.Move(-1, 1)) && free(elf.Move(0, 1)) && free(elf.Move(1, 1)) {
						delta = Coord{0, 1}
						break trying
					}
				case 2:
					if free(elf.Move(-1, -1)) && free(elf.Move(-1, 0)) && free(elf.Move(-1, 1)) {
						delta = Coord{-1, 0}
						break trying
					}
				case 3:
					if free(elf.Move(1, -1)) && free(elf.Move(1, 0)) && free(elf.Move(1, 1)) {
						delta = Coord{1, 0}
						break trying
					}
				}
			}
			if delta != Zero {
				destination := elf.Add(delta)
				if list, ok := destinations[destination]; ok {
					destinations[destination] = append(list, elfNo)
				} else {
					destinations[destination] = []int{elfNo}
				}
			}
		}

		for destination, sources := range destinations {
			if len(sources) == 1 {
				grove.elves[sources[0]] = destination
			}
		}

		if round == 10 {
			// Part 1
			ne, sw := grove.Bounds()
			fmt.Println((sw.x-ne.x+1)*(sw.y-ne.y+1) - len(grove.elves))
		}

		firstDirTry = (firstDirTry + 1) % 4
		round++
	}
}
