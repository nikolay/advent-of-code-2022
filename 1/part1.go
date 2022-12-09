package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	calories := int64(0)
	maxCalories := calories
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			if calories > maxCalories {
				maxCalories = calories
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
	if calories > maxCalories {
		maxCalories = calories
	}
	fmt.Println(maxCalories)
}
