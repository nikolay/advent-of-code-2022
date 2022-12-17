package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const ChamberWidth = 7

type Rock struct {
	width, height int
	sprite        []uint
}

type CacheKey struct {
	rockType, jet int
}

type CacheEntry struct {
	rockNumber, height int
}

func Solve(jets string, rocks []Rock, rocksToDrop int) int {
	var cave []uint
	cache := map[CacheKey]CacheEntry{}
	rockType := 0
	jet := 0
	height := 0
	for rockNumber := 0; rockNumber < rocksToDrop; rockNumber++ {
		rock := rocks[rockType]

		cacheKey := CacheKey{rockType, jet}
		if cacheEntry, ok := cache[cacheKey]; ok {
			rockLeftToDrop := rocksToDrop - rockNumber
			seriesLength := rockNumber - cacheEntry.rockNumber
			if rockLeftToDrop%seriesLength == 0 {
				return height + (height-cacheEntry.height)*rockLeftToDrop/seriesLength
			}
		}
		cache[cacheKey] = CacheEntry{rockNumber, height}

		rockX, rockY := 2, height+3+rock.height-1
		for rockY >= len(cave) {
			cave = append(cave, 0)
		}

		ok := true
		for ok {
			ch := jets[jet]
			jet = (jet + 1) % len(jets)
			var deltaX, deltaY int
			switch ch {
			case '<':
				deltaX, deltaY = -1, 0
			case '>':
				deltaX, deltaY = 1, 0
			}

			moveX, moveY := rockX+deltaX, rockY+deltaY
			if moveY >= rock.height-1 && moveX >= 0 && moveX <= ChamberWidth-rock.width {
				for spriteY := 0; ok && spriteY < rock.height; spriteY++ {
					bits := rock.sprite[spriteY] << moveX
					if cave[moveY-spriteY]^bits != cave[moveY-spriteY]|bits {
						ok = false
					}
				}
				if ok {
					rockX, rockY = moveX, moveY
				} else {
					ok = true
				}
			}

			deltaX, deltaY = 0, -1
			moveX, moveY = rockX+deltaX, rockY+deltaY
			if moveY < rock.height-1 || moveX < 0 || moveX > ChamberWidth-rock.width {
				ok = false
			} else {
				for spriteY := 0; ok && spriteY < rock.height; spriteY++ {
					bits := rock.sprite[spriteY] << moveX
					if cave[moveY-spriteY]^bits != cave[moveY-spriteY]|bits {
						ok = false
					}
				}
				if ok {
					rockX, rockY = moveX, moveY
				}
			}
		}

		for spriteY := 0; spriteY < rock.height; spriteY++ {
			cave[rockY-spriteY] |= rock.sprite[spriteY] << rockX
		}

		for height < len(cave) && cave[height] != 0 {
			height++
		}

		rockType = (rockType + 1) % len(rocks)
	}
	return height
}

func main() {
	// sprite bitmaps are mirrored horizontally
	rocks := []Rock{
		{
			4,
			1,
			[]uint{
				0b1111,
			},
		},
		{
			3,
			3,
			[]uint{
				0b010,
				0b111,
				0b010,
			},
		},
		{
			3,
			3,
			[]uint{
				0b100,
				0b100,
				0b111,
			},
		},
		{
			1,
			4,
			[]uint{
				0b1,
				0b1,
				0b1,
				0b1,
			},
		},
		{
			2,
			2,
			[]uint{
				0b11,
				0b11,
			},
		},
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		fmt.Println(Solve(line, rocks, 2022))
		fmt.Println(Solve(line, rocks, 1000000000000))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
