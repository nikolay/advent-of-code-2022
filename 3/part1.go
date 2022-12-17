package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Part1() int {
	const FullMask = 1<<2 - 1

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	score := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		l := len(line)
		if l == 0 {
			continue
		}
		h := l / 2
		bits := [Symbols]byte{}
		bit := byte(1)
		for i := 0; i < l; i++ {
			if i == h {
				bit <<= 1
			}
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
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return score
}
