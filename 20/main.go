package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Modulus(a, n int) int {
	remainder := a % n
	if remainder < 0 && n > 0 || remainder > 0 && n < 0 {
		return remainder + n
	}
	return remainder
}

func Mix(numbers []int, decryptionKey int, iterations int) int {
	count := len(numbers)
	decryptedNumbers := make([]int, count)
	indices := make([]int, count)
	for i, number := range numbers {
		decryptedNumbers[i] = number * decryptionKey
		indices[i] = i
	}
	for n := 0; n < iterations; n++ {
		for i := 0; i < count; i++ {
			oldPos := -1
			for pos := 0; pos < count; pos++ {
				if indices[pos] == i {
					oldPos = pos
					break
				}
			}
			newPos := Modulus(oldPos+decryptedNumbers[i], count-1)

			slice := append(indices[:oldPos], indices[oldPos+1:]...)
			mixed := make([]int, newPos+1, count)
			copy(mixed, slice[:newPos])
			mixed[newPos] = i
			mixed = append(mixed, slice[newPos:]...)

			indices = mixed
		}
	}
	for i := 0; i < count; i++ {
		if decryptedNumbers[i] == 0 {
			for pos := 0; pos < count; pos++ {
				if indices[pos] == i {
					sum := 0
					for _, offset := range []int{1000, 2000, 3000} {
						sum += decryptedNumbers[indices[Modulus(pos+offset, count)]]
					}
					return sum
				}
			}
		}
	}
	return -1
}

func Solve1(numbers []int) int {
	return Mix(numbers, 1, 1)
}

func Solve2(numbers []int) int {
	return Mix(numbers, 811589153, 10)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(Solve1(numbers))
	fmt.Println(Solve2(numbers))
}
