package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

func Part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := regexp.MustCompile(`^([ABC]) ([XYZ])$`)
	total := 0

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
		abc, xyz := matches[1], matches[2]
		score := 0
		switch xyz + abc {
		case "XA", "YC", "ZB":
			xyz = "Z"
		case "XB", "YA", "ZC":
			xyz = "X"
		case "XC", "YB", "ZA":
			xyz = "Y"
		}
		switch xyz {
		case "X":
			score += 1
		case "Y":
			score += 2
		case "Z":
			score += 3
		}
		switch abc + xyz {
		case "AX", "BY", "CZ":
			score += 3
		case "AZ", "BX", "CY":
		default:
			score += 6
		}
		total += score
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}
