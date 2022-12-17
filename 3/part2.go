package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Part2() int {
	const (
		Groups   = 3
		FullMask = 1<<Groups - 1
	)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	score := 0
	bits := [Symbols]byte{}
	elf := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		bit := byte(1) << elf
		for i := 0; i < len(line); i++ {
			c := line[i]
			var code int
			if c >= 'a' && c <= 'z' {
				code = int(c - 'a')
			} else if c >= 'A' && c <= 'Z' {
				code = Letters + int(c-'A')
			} else {
				log.Fatalf("invalid character: '%c'", c)
			}
			bits[code] |= bit
			if bits[code] == FullMask {
				score += 1 + code
				break
			}
		}
		elf++
		if elf == Groups {
			elf = 0
			bits = [Symbols]byte{}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return score
}
