package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part2() string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	heap := map[int]string{}
	var keys []string
	keymap := map[string]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		for i, buf := 0, line; len(buf) > 0; i++ {
			str := strings.TrimSpace(buf[:3])
			buf = buf[3:]
			if len(buf) > 0 {
				if buf[0] != ' ' {
					log.Fatalf("invalid command: %v", line)
				}
				buf = buf[1:]
			}
			if len(str) == 0 {
				continue
			}
			if str[0] == '[' && str[2] == ']' {
				contents := str[1:2]
				if v, ok := heap[i]; ok {
					heap[i] = v + contents
				} else {
					heap[i] = contents
				}
			} else {
				keys = append(keys, str)
				keymap[str] = i
			}
		}
	}

	r := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		matches := r.FindStringSubmatch(line)
		if len(matches) == 0 {
			log.Fatalf("invalid command: %v", line)
		}
		count, _ := strconv.Atoi(matches[1])
		from := keymap[matches[2]]
		to := keymap[matches[3]]

		str := heap[from][:count]
		heap[from] = heap[from][count:]
		heap[to] = str + heap[to]
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := ""
	for _, key := range keys {
		result += heap[keymap[key]][:1]
	}
	return result
}
