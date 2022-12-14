package main

func Part2(input *Input) (result uint) {
	space := input.GetSpace()
	for Drop(&space, input.drop, 0, input.maxY+2) != BLOCKED {
		result++
	}
	return
}
