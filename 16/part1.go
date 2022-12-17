package main

type State1 struct {
	valve      int
	openValves uint64
	pressure   int
}

type Key1 struct {
	valve      int
	openValves uint64
}

func Part1(input *Input, time int) int {
	states := []State1{{valve: input.FindValve("AA"), openValves: 0, pressure: 0}}
	best := map[Key1]int{}
	for t := 1; t < time; t++ {
		var newStates []State1
		for _, state := range states {
			key := Key1{state.valve, state.openValves}
			if bestPressure, ok := best[key]; ok && state.pressure <= bestPressure {
				continue
			}
			best[key] = state.pressure

			valve, _ := input.valves[state.valve]
			flowRate := valve.flowRate
			bitmask := uint64(1) << state.valve
			if state.openValves&bitmask == 0 && flowRate > 0 {
				newStates = append(
					newStates,
					State1{state.valve, state.openValves | bitmask, state.pressure + (time-t)*flowRate},
				)
			}

			for _, tunnel := range valve.tunnels {
				newStates = append(
					newStates,
					State1{tunnel, state.openValves, state.pressure})
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
