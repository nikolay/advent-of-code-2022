package main

type State2 struct {
	me         int
	elephant   int
	openValves uint
	pressure   int
}

type CacheKey2 struct {
	low, high  int
	openValves uint
}

func Order(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func Part2(input *Input, time int) int {
	index := input.FindValve("AA")
	states := []State2{{me: index, elephant: index, openValves: 0, pressure: 0}}
	best := map[CacheKey2]int{}
	for t := 1; t < time; t++ {
		var newStates []State2
		for _, state := range states {
			low, high := Order(state.me, state.elephant)
			key := CacheKey2{low, high, state.openValves}
			if bestPressure, ok := best[key]; ok && state.pressure <= bestPressure {
				continue
			}
			best[key] = state.pressure

			me, _ := input.valves[state.me]
			elephant, _ := input.valves[state.elephant]

			flowRate := me.flowRate
			elephantFlowRate := elephant.flowRate

			bitmask := uint(1) << state.me
			elephantBitmask := uint(1) << state.elephant

			open := flowRate > 0 && state.openValves&bitmask == 0
			elephantOpen := elephantFlowRate > 0 && state.openValves&elephantBitmask == 0

			if open {
				if elephantOpen {
					newStates = append(
						newStates,
						State2{state.me, state.elephant, state.openValves | bitmask | elephantBitmask, state.pressure + (time-t)*(flowRate+elephantFlowRate)},
					)
				}
				for _, tunnel := range elephant.tunnels {
					if tunnel != state.me {
						newStates = append(
							newStates,
							State2{state.me, tunnel, state.openValves | bitmask, state.pressure + (time-t)*flowRate},
						)
					}
				}
			}

			if elephantOpen {
				for _, tunnel := range me.tunnels {
					if tunnel != state.elephant {
						newStates = append(
							newStates,
							State2{tunnel, state.elephant, state.openValves | elephantBitmask, state.pressure + (time-t)*elephantFlowRate},
						)
					}
				}
			}

			var trail [64][64]bool
			trail[state.me][state.elephant] = true
			trail[state.elephant][state.me] = true
			for _, tunnel := range me.tunnels {
				for _, elephantTunnel := range elephant.tunnels {
					if elephantTunnel == tunnel || trail[tunnel][elephantTunnel] {
						continue
					}
					newStates = append(
						newStates,
						State2{tunnel, elephantTunnel, state.openValves, state.pressure},
					)
					trail[tunnel][elephantTunnel] = true
					trail[elephantTunnel][tunnel] = true
				}
			}

			states = newStates
		}
	}
	maxPressure := 0
	for _, state := range states {
		if state.pressure > maxPressure {
			maxPressure = state.pressure
		}
	}
	return maxPressure
}
