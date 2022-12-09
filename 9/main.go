package main

type Position struct {
	x, y int
}

func (pos Position) Move(dx, dy int) Position {
	return Position{pos.x + dx, pos.y + dy}
}

func (pos Position) Distance(p Position) Position {
	return Position{pos.x - p.x, pos.y - p.y}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return +1
	}
	return 0
}

func DirDelta(dir string) Position {
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
	part1()
	part2()
}
