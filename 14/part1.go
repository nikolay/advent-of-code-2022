package main

func Part1(input *Input) (result int) {
	space := input.GetSpace()

	for Drop(&space, input.drop, input.maxY+1) {
		result++
	}
	return
}
