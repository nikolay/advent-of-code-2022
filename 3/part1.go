package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	FULL_MASK = 1<<2 - 1
	LETTERS   = 26
	SYMBOLS   = 2 * LETTERS
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	score := uint64(0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		l := len(line)
		if l == 0 {
			continue
		}
		h := l / 2
		bits := [SYMBOLS]byte{}
		bit := byte(1)
		for i := 0; i < l; i++ {
			if i == h {
				bit <<= 1
			}
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
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(score)
}
