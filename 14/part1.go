package main

func Part1(input *Input) (result uint) {
	space := input.GetSpace()
	for Drop(&space, input.drop, input.maxY+1, 0) == LANDED {
		result++
	}
	return
}
