package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Part2() {
	const (
		Groups   = 3
		FullMask = 1<<Groups - 1
	)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	score := uint64(0)
	bits := [Symbols]byte{}
	elf := uint(0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		bit := byte(1) << elf
		for i := 0; i < len(line); i++ {
			c := line[i]
			var code uint
			if c >= 'a' && c <= 'z' {
				code = uint(c - 'a')
			} else if c >= 'A' && c <= 'Z' {
				code = Letters + uint(c-'A')
			} else {
				log.Fatalf("invalid character: '%c'", c)
			}
			bits[code] |= bit
			if bits[code] == FullMask {
				score += 1 + uint64(code)
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
	fmt.Println(score)
}
