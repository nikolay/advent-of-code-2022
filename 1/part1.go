package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	calories := int64(0)
	max_calories := calories
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			if calories > max_calories {
				max_calories = calories
			}
			calories = 0
			continue
		}
		if calorie, err := strconv.ParseInt(line, 10, 64); err != nil {
			log.Fatal(err)
		} else {
			calories += calorie
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if calories > max_calories {
		max_calories = calories
	}
	fmt.Println(max_calories)
}
