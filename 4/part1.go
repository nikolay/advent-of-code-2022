package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)
	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		matches := r.FindStringSubmatch(line)
		if len(matches) == 0 {
			log.Fatalf("invalid command: %v", line)
		}
		from1, _ := strconv.Atoi(matches[1])
		to1, _ := strconv.Atoi(matches[2])
		from2, _ := strconv.Atoi(matches[3])
		to2, _ := strconv.Atoi(matches[4])
		if from1 >= from2 && to1 <= to2 || from2 >= from1 && to2 <= to1 {
			count++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return count
}
