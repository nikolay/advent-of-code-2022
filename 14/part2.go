package main

func Part2(input *Input) (result int) {
	space := input.GetSpace()
	for Drop(&space, input.drop, 0, input.maxY+2) != BLOCKED {
		result++
	}
	return
}
