package main

type Position struct {
	x, y int
}

func (pos Position) Move(dx, dy int) Position {
	return Position{pos.x + dx, pos.y + dy}
}

func (pos Position) Add(p Position) Position {
	return Position{pos.x + p.x, pos.y + p.y}
}

func (pos Position) Subtract(p Position) Position {
	return Position{pos.x - p.x, pos.y - p.y}
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
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

func DirPosition(dir string) Position {
	switch dir {
	case "L":
		return Position{-1, 0}
	case "R":
		return Position{+1, 0}
	case "U":
		return Position{0, -1}
	case "D":
		return Position{0, +1}
	}
	return Position{0, 0}
}

func main() {
	Part1()
	Part2()
}
