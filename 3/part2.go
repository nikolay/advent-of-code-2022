package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	GROUPS    = 3
	FULL_MASK = 1<<GROUPS - 1
	LETTERS   = 26
	SYMBOLS   = 2 * LETTERS
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	score := uint64(0)
	bits := [SYMBOLS]byte{}
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
				code = LETTERS + uint(c-'A')
			} else {
				log.Fatalf("invalid character: '%c'", c)
			}
			bits[code] |= bit
			if bits[code] == FULL_MASK {
				score += 1 + uint64(code)
				break
			}
		}
		elf++
		if elf == GROUPS {
			elf = 0
			bits = [SYMBOLS]byte{}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(score)
}
