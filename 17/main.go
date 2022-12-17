package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Rock struct {
	width, height int64
	sprite        []byte
}

type Key struct {
	rock, jet int
}

type CacheEntry struct {
	rockNumber, height int64
}

func Solve(jets string, rocks []Rock, rocksToDrop int64) int64 {
	var chamber []byte
	cache := map[Key]CacheEntry{}
	fallingRock := 0
	jet := 0
	deletedRows := int64(0)
	height := int64(0)
	for rockNumber := int64(0); rockNumber < rocksToDrop; rockNumber++ {
		rock := rocks[fallingRock]

	height:
		for height < int64(len(chamber)) {
			switch chamber[height] {
			case 0:
				break height
			case 0b1111111:
				chamber = chamber[height+1:]
				deletedRows += int64(height)
				height = 0
				continue
			}
			height++
		}

		actualHeight := deletedRows + height
		cacheKey := Key{fallingRock, jet}
		if cacheEntry, ok := cache[cacheKey]; ok {
			rockLeftToDrop := rocksToDrop - rockNumber
			seriesLength := rockNumber - cacheEntry.rockNumber
			if rockLeftToDrop%seriesLength == 0 {
				return actualHeight + (actualHeight-cacheEntry.height)*rockLeftToDrop/seriesLength
			}
		}
		cache[cacheKey] = CacheEntry{rockNumber, actualHeight}

		rockX, rockY := int64(2), height+3+rock.height-1
		for rockY >= int64(len(chamber)) {
			chamber = append(chamber, 0b0000000)
		}

		ok := true
		for ok {
			ch := jets[jet]
			jet = (jet + 1) % len(jets)
			var deltaX, deltaY int64
			switch ch {
			case '<':
				deltaX, deltaY = -1, 0
			case '>':
				deltaX, deltaY = 1, 0
			}

			moveX, moveY := rockX+deltaX, rockY+deltaY
			if moveY >= rock.height-1 && moveX >= 0 && moveX <= 7-rock.width {
				for spriteY := int64(0); ok && spriteY < rock.height; spriteY++ {
					bits := rock.sprite[spriteY] << moveX
					if chamber[moveY-spriteY]^bits != chamber[moveY-spriteY]|bits {
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
			if moveY < rock.height-1 || moveX < 0 || moveX > 7-rock.width {
				ok = false
			} else {
				for spriteY := int64(0); ok && spriteY < rock.height; spriteY++ {
					bits := rock.sprite[spriteY] << moveX
					if chamber[moveY-spriteY]^bits != chamber[moveY-spriteY]|bits {
						ok = false
					}
				}
				if ok {
					rockX, rockY = moveX, moveY
				}
			}
		}

		for spriteY := int64(0); spriteY < rock.height; spriteY++ {
			chamber[rockY-spriteY] |= rock.sprite[spriteY] << rockX
		}

		fallingRock = (fallingRock + 1) % len(rocks)
	}

	var y int64
	for y = height; chamber[y] != 0; y++ {
	}
	return deletedRows + y
}

func main() {
	// sprite bitmaps are mirrored horizontally
	rocks := []Rock{
		{
			4,
			1,
			[]byte{
				0b1111,
			},
		},
		{
			3,
			3,
			[]byte{
				0b010,
				0b111,
				0b010,
			},
		},
		{
			3,
			3,
			[]byte{
				0b100,
				0b100,
				0b111,
			},
		},
		{
			1,
			4,
			[]byte{
				0b1,
				0b1,
				0b1,
				0b1,
			},
		},
		{
			2,
			2,
			[]byte{
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
