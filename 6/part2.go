package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const STREAK_SIZE = 14

func main() {
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
		for i := 0; i < len(line)-STREAK_SIZE; i++ {
			for j := i; j < i+STREAK_SIZE-1; j++ {
				for k := j + 1; k < i+STREAK_SIZE; k++ {
					if line[j] == line[k] {
						continue outer
					}
				}
			}
			fmt.Println(i + STREAK_SIZE)
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
