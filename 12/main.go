package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
	row, col int
}

func Height(ch byte) int {
	switch ch {
	case 'S':
		ch = 'a'
	case 'E':
		ch = 'z'
	}
	return int(ch) - int('a')
}

func Wave(rows []string, starts []Coord, end Coord) int {
	moves := []Coord{
		Coord{-1, 0},
		Coord{0, +1},
		Coord{+1, 0},
		Coord{0, -1},
	}
	height, width := len(rows), len(rows[0])
	wave := [][]int{}
	for row := 0; row < height; row++ {
		r := make([]int, width)
		for col := 0; col < width; col++ {
			r[col] = -1
		}
		wave = append(wave, r)
	}
	step := 0
	for _, s := range starts {
		wave[s.row][s.col] = step
	}
	for {
		found := false
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				h := Height(rows[row][col])
				if wave[row][col] == step {
					for _, move := range moves {
						r := row + move.row
						c := col + move.col
						if r >= 0 && r < height && c >= 0 && c < width && wave[r][c] == -1 {
							if Height(rows[r][c]) <= h+1 {
								wave[r][c] = step + 1
								found = true
							}
						}
					}
				}
			}
		}
		if !found {
			break
		}
		if finish := wave[end.row][end.col]; finish >= 0 {
			return finish
		}
		step++
	}
	return -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rows := []string{}
	start, end := Coord{0, 0}, Coord{0, 0}
	for row := 0; scanner.Scan(); row++ {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		rows = append(rows, line)
		s := strings.Index(line, "S")
		if s >= 0 {
			start = Coord{row, s}
		}
		e := strings.Index(line, "E")
		if e >= 0 {
			end = Coord{row, e}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	starts := []Coord{start}
	fmt.Println(Wave(rows, starts, end))

	// Part 2
	for row := 0; row < len(rows); row++ {
		for col := 0; col < len(rows[row]); col++ {
			if rows[row][col] == 'a' {
				starts = append(starts, Coord{row, col})
			}
		}
	}
	fmt.Println(Wave(rows, starts, end))
}
