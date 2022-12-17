package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	calories := 0
	maxCalories1 := calories
	maxCalories2 := calories
	maxCalories3 := calories

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			if calories > maxCalories1 {
				maxCalories3 = maxCalories2
				maxCalories2 = maxCalories1
				maxCalories1 = calories
			} else if calories > maxCalories2 {
				maxCalories3 = maxCalories2
				maxCalories2 = calories
			} else if calories > maxCalories3 {
				maxCalories3 = calories
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

	if calories > maxCalories1 {
		maxCalories3 = maxCalories2
		maxCalories2 = maxCalories1
		maxCalories1 = calories
	} else if calories > maxCalories2 {
		maxCalories3 = maxCalories2
		maxCalories2 = calories
	} else if calories > maxCalories3 {
		maxCalories3 = calories
	}
	return maxCalories1 + maxCalories2 + maxCalories3
}
