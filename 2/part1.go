package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(`^([ABC]) ([XYZ])$`)
	total := int64(0)
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
		score := int64(0)
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
	fmt.Println(total)
}
