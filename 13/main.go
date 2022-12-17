package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	hasValue  bool
	value     int
	list      []Node
	isDivider bool
}

func Sign(n int) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return +1
	}
	return 0
}

func Parse(line string, startPos int) (Node, int, error) {
	pos := startPos
	if line[pos] == '[' {
		pos++
		result := Node{}
		result.list = []Node{}
		for line[pos] != ']' {
			node, newPos, err := Parse(line, pos)
			if err != nil {
				return Node{}, 0, err
			}
			result.list = append(result.list, node)
			pos = newPos
			if ch := line[pos]; ch == ',' {
				pos++
			} else if ch != ']' {
				return Node{}, 0, fmt.Errorf("invalid command: %v", line)
			}
		}
		return result, pos + 1, nil
	}
	result := Node{hasValue: true}
	for line[pos] >= '0' && line[pos] <= '9' {
		pos++
	}
	var err error
	if result.value, err = strconv.Atoi(line[startPos:pos]); err != nil {
		return Node{}, 0, err
	}
	return result, pos, nil
}

func Compare(left, right Node) int {
	if left.hasValue {
		if right.hasValue {
			return Sign(left.value - right.value)
		}
		return Compare(Node{list: []Node{left}}, right)
	}
	if right.hasValue {
		return Compare(left, Node{list: []Node{right}})
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

	var nodes []Node

	scanner := bufio.NewScanner(file)
	for row := 0; scanner.Scan(); row++ {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if node, _, err := Parse(line, 0); err != nil {
			log.Fatal(err)
		} else {
			nodes = append(nodes, node)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	sum := 0
	for i := 0; i < len(nodes); i += 2 {
		if Compare(nodes[i], nodes[i+1]) < 0 {
			sum += i/2 + 1
		}
	}
	fmt.Println(sum)

	// Part 2
	var divisors []Node
	var indices []int
	for index, v := range []int{2, 6} {
		divisors = append(divisors, Node{isDivider: true, list: []Node{{list: []Node{{hasValue: true, value: v}}}}})
		indices = append(indices, index)
	}
	for _, node := range nodes {
		for index, divisor := range divisors {
			if Compare(node, divisor) < 0 {
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
