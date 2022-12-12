package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	calories := int64(0)
	maxCalories1 := calories
	maxCalories2 := calories
	maxCalories3 := calories
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
		if calorie, err := strconv.ParseInt(line, 10, 64); err != nil {
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
	fmt.Println(maxCalories1 + maxCalories2 + maxCalories3)
}
