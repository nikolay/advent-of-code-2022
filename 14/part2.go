package main

func Part2(input *Input) (result int) {
	space := input.GetSpace()

	y := input.maxY + 2
	offset := y - input.drop.y
	for x := input.drop.x - offset; x <= input.drop.x+offset; x++ {
		space.Put(Coord{x, y}, '#')
	}

	for Drop(&space, input.drop, y) {
		result++
	}
	return
}
