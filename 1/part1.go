package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	calories := 0
	maxCalories := calories

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			if calories > maxCalories {
				maxCalories = calories
			}
			calories = 0
			continue
		}
		if calorie, err := strconv.Atoi(line); err != nil {
			log.Fatal(err)
		} else {
			calories += calorie
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if calories > maxCalories {
		maxCalories = calories
	}
	return maxCalories
}
