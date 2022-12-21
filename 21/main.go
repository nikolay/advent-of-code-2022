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
	Literal Operator = iota
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
	return Literal, fmt.Errorf("invalid operator: %v", op)
}

type Side struct {
	isLiteral bool
	literal   int
	monkey    string
}

func ParseSide(side string) Side {
	if literal, err := strconv.Atoi(side); err == nil {
		return Side{true, literal, ""}
	}
	return Side{false, 0, side}
}

type Monkey struct {
	name        string
	op          Operator
	left, right Side
}

func Compute(values *map[string]Monkey) (bool, int) {
	for {
		var deleteMonkeys []string
		for yellingMonkeyName, yellingMonkey := range *values {
			if yellingMonkey.op == Literal {
				for waitingMonkeyName, waitingMonkey := range *values {
					if waitingMonkey.op != Literal {
						if !waitingMonkey.left.isLiteral && waitingMonkey.left.monkey == yellingMonkeyName {
							waitingMonkey.left = yellingMonkey.left
						}
						if !waitingMonkey.right.isLiteral && waitingMonkey.right.monkey == yellingMonkeyName {
							waitingMonkey.right = yellingMonkey.left
						}
						if waitingMonkey.left.isLiteral && waitingMonkey.right.isLiteral {
							var literal int
							switch waitingMonkey.op {
							case Add:
								literal = waitingMonkey.left.literal + waitingMonkey.right.literal
							case Subtract:
								literal = waitingMonkey.left.literal - waitingMonkey.right.literal
							case Multiply:
								literal = waitingMonkey.left.literal * waitingMonkey.right.literal
							case Divide:
								literal = waitingMonkey.left.literal / waitingMonkey.right.literal
							}
							if waitingMonkeyName == "root" {
								return true, literal
							}
							waitingMonkey.op = Literal
							waitingMonkey.left = Side{true, literal, ""}
							waitingMonkey.right = Side{}
						}
						(*values)[waitingMonkeyName] = waitingMonkey
					}
				}
				deleteMonkeys = append(deleteMonkeys, yellingMonkeyName)
			}
		}
		if len(deleteMonkeys) == 0 {
			break
		}
		for _, name := range deleteMonkeys {
			delete(*values, name)
		}
	}
	return false, 0
}

func Solve1(monkeys map[string]Monkey) int {
	values := map[string]Monkey{}
	for name, monkey := range monkeys {
		values[name] = monkey
	}
	found, result := Compute(&values)
	if !found {
		log.Fatalf("no value for root found")
	}
	return result
}

func Solve2(monkeys map[string]Monkey) int {
	values := map[string]Monkey{}
	for name, monkey := range monkeys {
		if name == "humn" || name == "root" {
			continue
		}
		values[name] = monkey
	}

	root, _ := monkeys["root"]
	values["root"] = Monkey{"root", Subtract, root.left, root.right}

	Compute(&values)

	inferred := map[string]int{}

	root = values["root"]
	if root.left.isLiteral {
		inferred[root.right.monkey] = root.left.literal
	} else if root.right.isLiteral {
		inferred[root.left.monkey] = root.right.literal
	} else {
		log.Fatalf("could not find a solution")
	}

	for {
		any := false
		for name, value := range inferred {
			monkey := values[name]
			var reversedValue int
			if monkey.left.isLiteral {
				switch monkey.op {
				case Add:
					reversedValue = value - monkey.left.literal
				case Subtract:
					reversedValue = monkey.left.literal - value
				case Multiply:
					reversedValue = value / monkey.left.literal
				case Divide:
					reversedValue = monkey.left.literal / value
				}
				if monkey.right.monkey == "humn" {
					return reversedValue
				}
				inferred[monkey.right.monkey] = reversedValue
				any = true
			} else if monkey.right.isLiteral {
				switch monkey.op {
				case Add:
					reversedValue = value - monkey.right.literal
				case Subtract:
					reversedValue = monkey.right.literal + value
				case Multiply:
					reversedValue = value / monkey.right.literal
				case Divide:
					reversedValue = monkey.right.literal * value
				}
				if monkey.left.monkey == "humn" {
					return reversedValue
				}
				inferred[monkey.left.monkey] = reversedValue
				any = true
			}
		}
		if !any {
			break
		}
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
			monkey = Monkey{name, Literal, ParseSide(expr[0]), Side{}}
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
