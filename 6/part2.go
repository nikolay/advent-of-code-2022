package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Part2() {
	const StreakSize = 14

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
	outer:
		for i := 0; i < len(line)-StreakSize; i++ {
			for j := i; j < i+StreakSize-1; j++ {
				for k := j + 1; k < i+StreakSize; k++ {
					if line[j] == line[k] {
						continue outer
					}
				}
			}
			fmt.Println(i + StreakSize)
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
