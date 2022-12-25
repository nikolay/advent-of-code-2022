package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
	x, y int
}

var Zero = Coord{}

func (coord Coord) Move(dir int) Coord {
	return coord.Add([]Coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}[dir])
}

func (coord Coord) Add(add Coord) Coord {
	return Coord{coord.x + add.x, coord.y + add.y}
}

func SnafuToDecimal(snafu string) int {
	decimal := 0
	for _, ch := range snafu {
		var digit int
		switch ch {
		case '=':
			digit = -2
		case '-':
			digit = -1
		default:
			digit = int(ch - '0')
		}
		decimal *= 5
		decimal += digit
	}
	return decimal
}

func DecimalToSnafu(decimal int) string {
	switch decimal {
	case -2:
		return "="
	case -1:
		return "-"
	case 0:
		return "0"
	}
	snafu := ""
	for decimal > 0 {
		digit := decimal % 5
		if digit > 2 {
			decimal += 5
		}
		snafu = []string{"0", "1", "2", "=", "-"}[digit] + snafu
		decimal = decimal / 5
	}
	return snafu
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		sum += SnafuToDecimal(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(DecimalToSnafu(sum))
}
