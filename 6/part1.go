package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Part1() int {
	const StreakSize = 4

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
		for i := 0; i < len(line)-StreakSize; i++ {
			if line[i] != line[i+1] && line[i] != line[i+2] && line[i] != line[i+3] &&
				line[i+1] != line[i+2] && line[i+1] != line[i+3] &&
				line[i+2] != line[i+3] {
				return i + StreakSize
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return Error
}
