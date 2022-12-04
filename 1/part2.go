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
	max_calories_1 := calories
	max_calories_2 := calories
	max_calories_3 := calories
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			if calories > max_calories_1 {
				max_calories_3 = max_calories_2
				max_calories_2 = max_calories_1
				max_calories_1 = calories
			} else if calories > max_calories_2 {
				max_calories_3 = max_calories_2
				max_calories_2 = calories
			} else if calories > max_calories_3 {
				max_calories_3 = calories
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
	if calories > max_calories_1 {
		max_calories_3 = max_calories_2
		max_calories_2 = max_calories_1
		max_calories_1 = calories
	} else if calories > max_calories_2 {
		max_calories_3 = max_calories_2
		max_calories_2 = calories
	} else if calories > max_calories_3 {
		max_calories_3 = calories
	}
	fmt.Println(max_calories_1 + max_calories_2 + max_calories_3)
}
