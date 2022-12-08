package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const STREAK_SIZE = 4

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
		for i := 0; i < len(line)-STREAK_SIZE; i++ {
			if line[i] != line[i+1] && line[i] != line[i+2] && line[i] != line[i+3] &&
				line[i+1] != line[i+2] && line[i+1] != line[i+3] &&
				line[i+2] != line[i+3] {
				fmt.Println(i + STREAK_SIZE)
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
