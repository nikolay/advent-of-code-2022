package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operator int

const (
	Nop Operator = iota
	Add
	Subtract
	Multiply
	Divide
)

func ParseOperator(op string) (Operator, error) {
	switch op {
	case "+":
		return Add, nil
	case "-":
		return Subtract, nil
	case "*":
		return Multiply, nil
	case "/":
		return Divide, nil
	}
	return Nop, fmt.Errorf("invalid operator: %v", op)
}

type Side struct {
	yelled bool
	number int
	monkey string
}

func ParseSide(side string) Side {
	if number, err := strconv.Atoi(side); err == nil {
		return Side{true, number, ""}
	}
	return Side{false, 0, side}
}

type Monkey struct {
	name        string
	op          Operator
	left, right Side
}

func Compute(state *map[string]Monkey) (bool, int) {
	for {
		var monkeysToRetire []string
		for yellingMonkeyName, yellingMonkey := range *state {
			if yellingMonkey.op == Nop {
				for waitingMonkeyName, waitingMonkey := range *state {
					if waitingMonkey.op != Nop {
						if !waitingMonkey.left.yelled && waitingMonkey.left.monkey == yellingMonkeyName {
							waitingMonkey.left = yellingMonkey.left
						}
						if !waitingMonkey.right.yelled && waitingMonkey.right.monkey == yellingMonkeyName {
							waitingMonkey.right = yellingMonkey.left
						}
						if waitingMonkey.left.yelled && waitingMonkey.right.yelled {
							var number int
							switch waitingMonkey.op {
							case Add:
								number = waitingMonkey.left.number + waitingMonkey.right.number
							case Subtract:
								number = waitingMonkey.left.number - waitingMonkey.right.number
							case Multiply:
								number = waitingMonkey.left.number * waitingMonkey.right.number
							case Divide:
								number = waitingMonkey.left.number / waitingMonkey.right.number
							}
							if waitingMonkeyName == "root" {
								return true, number
							}
							waitingMonkey.op = Nop
							waitingMonkey.left = Side{true, number, ""}
							waitingMonkey.right = Side{}
						}
						(*state)[waitingMonkeyName] = waitingMonkey
					}
				}
				monkeysToRetire = append(monkeysToRetire, yellingMonkeyName)
			}
		}
		if len(monkeysToRetire) == 0 {
			break
		}
		for _, monkeyName := range monkeysToRetire {
			delete(*state, monkeyName)
		}
	}
	return false, 0
}

func Solve1(monkeys map[string]Monkey) int {
	state := map[string]Monkey{}
	for monkeyName, monkey := range monkeys {
		state[monkeyName] = monkey
	}
	rootYelled, rootNumber := Compute(&state)
	if !rootYelled {
		log.Fatalf("root couldn't yell a number")
	}
	return rootNumber
}

func Solve2(monkeys map[string]Monkey) int {
	state := map[string]Monkey{}
	for monkeyName, monkey := range monkeys {
		if monkeyName == "humn" || monkeyName == "root" {
			continue
		}
		state[monkeyName] = monkey
	}

	root, _ := monkeys["root"]
	state["root"] = Monkey{"root", Subtract, root.left, root.right}

	Compute(&state)

	inferredNumbers := map[string]int{}

	root = state["root"]
	if root.left.yelled {
		inferredNumbers[root.right.monkey] = root.left.number
	} else if root.right.yelled {
		inferredNumbers[root.left.monkey] = root.right.number
	} else {
		log.Fatalf("could not find a solution")
	}

	for {
		newInferredNumbers := map[string]int{}
		found := 0
		for monkeyName, number := range inferredNumbers {
			monkey := state[monkeyName]
			var inferredNUmber int
			if monkey.left.yelled {
				switch monkey.op {
				case Add:
					inferredNUmber = number - monkey.left.number
				case Subtract:
					inferredNUmber = monkey.left.number - number
				case Multiply:
					inferredNUmber = number / monkey.left.number
				case Divide:
					inferredNUmber = monkey.left.number / number
				}
				if monkey.right.monkey == "humn" {
					return inferredNUmber
				}
				newInferredNumbers[monkey.right.monkey] = inferredNUmber
				found++
			} else if monkey.right.yelled {
				switch monkey.op {
				case Add:
					inferredNUmber = number - monkey.right.number
				case Subtract:
					inferredNUmber = monkey.right.number + number
				case Multiply:
					inferredNUmber = number / monkey.right.number
				case Divide:
					inferredNUmber = monkey.right.number * number
				}
				if monkey.left.monkey == "humn" {
					return inferredNUmber
				}
				newInferredNumbers[monkey.left.monkey] = inferredNUmber
				found++
			}
		}
		if found == 0 {
			break
		}
		inferredNumbers = newInferredNumbers
	}

	return -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	monkeys := map[string]Monkey{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		colon := strings.Index(line, ":")
		name := line[0:colon]
		expr := strings.Fields(line[colon+1:])

		var monkey Monkey
		if len(expr) == 1 {
			monkey = Monkey{name, Nop, ParseSide(expr[0]), Side{}}
		} else {
			op, err := ParseOperator(expr[1])
			if err != nil {
				log.Fatal(err)
			}
			monkey = Monkey{name, op, ParseSide(expr[0]), ParseSide(expr[2])}
		}
		monkeys[name] = monkey
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(Solve1(monkeys))
	fmt.Println(Solve2(monkeys))
}
