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

	var rows [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		var row []int
		for i := 0; i < len(line); i++ {
			height, _ := strconv.Atoi(line[i : i+1])
			row = append(row, height)
		}
		rows = append(rows, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	height, width := len(rows), len(rows[0])
	solution1 := 2 * (width + (height - 2))
	solution2 := 0
	for row := 1; row < height-1; row++ {
		for col := 1; col < width-1; col++ {
			h := rows[row][col]
			views := 4
			top := 0
			for i := row - 1; i >= 0; i-- {
				top++
				if rows[i][col] >= h {
					views--
					break
				}
			}
			bottom := 0
			for i := row + 1; i < height; i++ {
				bottom++
				if rows[i][col] >= h {
					views--
					break
				}
			}
			left := 0
			for i := col - 1; i >= 0; i-- {
				left++
				if rows[row][i] >= h {
					views--
					break
				}
			}
			right := 0
			for i := col + 1; i < width; i++ {
				right++
				if rows[row][i] >= h {
					views--
					break
				}
			}
			if views > 0 {
				solution1++
			}
			if score := top * bottom * left * right; score > solution2 {
				solution2 = score
			}
		}
	}
	fmt.Println(solution1)
	fmt.Println(solution2)
}
