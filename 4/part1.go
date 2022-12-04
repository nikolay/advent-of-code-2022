package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)
	count := int64(0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		matches := r.FindStringSubmatch(line)
		if len(matches) == 0 {
			log.Fatalf("invalid command: %v", line)
		}
		f1, _ := strconv.Atoi(matches[1])
		t1, _ := strconv.Atoi(matches[2])
		f2, _ := strconv.Atoi(matches[3])
		t2, _ := strconv.Atoi(matches[4])
		if f1 >= f2 && t1 <= t2 || f2 >= f1 && t2 <= t1 {
			count++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
