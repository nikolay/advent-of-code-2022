package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	hasValue  bool
	value     int
	list      []Item
	isDivider bool
}

func Parse(line string, startPos int) (Item, int, error) {
	pos := startPos
	if line[pos] == '[' {
		pos++
		result := Item{}
		result.list = []Item{}
		for line[pos] != ']' {
			item, newPos, err := Parse(line, pos)
			if err != nil {
				return Item{}, 0, err
			}
			result.list = append(result.list, item)
			pos = newPos
			if ch := line[pos]; ch == ',' {
				pos++
			} else if ch != ']' {
				return Item{}, 0, fmt.Errorf("invalid command: %v", line)
			}
		}
		return result, pos + 1, nil
	}
	result := Item{hasValue: true}
	for line[pos] >= '0' && line[pos] <= '9' {
		pos++
	}
	var err error
	if result.value, err = strconv.Atoi(line[startPos:pos]); err != nil {
		return Item{}, 0, err
	}
	return result, pos, nil
}

func Sign(v int) int {
	if v < 0 {
		return -1
	}
	if v > 0 {
		return +1
	}
	return 0
}

func Compare(left, right Item) int {
	if left.hasValue {
		if right.hasValue {
			return Sign(left.value - right.value)
		}
		return Compare(Item{list: []Item{left}}, right)
	}
	if right.hasValue {
		return Compare(left, Item{list: []Item{right}})
	}
	l, r := len(left.list), len(right.list)
	for i := 0; i < l && i < r; i++ {
		if v := Compare(left.list[i], right.list[i]); v != 0 {
			return v
		}
	}
	return Sign(l - r)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	items := []Item{}
	for row := 0; scanner.Scan(); row++ {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if item, _, err := Parse(line, 0); err != nil {
			log.Fatal(err)
		} else {
			items = append(items, item)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	sum := 0
	for i := 0; i < len(items); i += 2 {
		if Compare(items[i], items[i+1]) < 0 {
			sum += i/2 + 1
		}
	}
	fmt.Println(sum)

	// Part 2
	divisors, indices := []Item{}, []int{}
	for index, v := range []int{2, 6} {
		divisors = append(divisors, Item{isDivider: true, list: []Item{Item{list: []Item{Item{hasValue: true, value: v}}}}})
		indices = append(indices, index)
	}
	for _, item := range items {
		for index, divisor := range divisors {
			if Compare(item, divisor) < 0 {
				indices[index]++
			}
		}
	}
	product := 1
	for _, index := range indices {
		product *= index + 1
	}
	fmt.Println(product)
}
