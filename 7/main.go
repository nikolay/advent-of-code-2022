package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var dirs map[string]uint64
	var pwd []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		fields := strings.Fields(line)
		if fields[0] == "$" {
			if fields[1] == "cd" {
				dir := fields[2]
				switch dir {
				case "..":
					if len(pwd) > 0 {
						pwd = pwd[:len(pwd)-1]
					}
				case "/":
					pwd = []string{}
				default:
					pwd = append(pwd, dir)
				}
			}
		} else if fields[0] != "dir" {
			filesize, _ := strconv.ParseUint(fields[0], 10, 64)
			for n := len(pwd); n >= 0; n-- {
				path := strings.Join(pwd[0:n], "/")
				dirsize := filesize
				if v, ok := dirs[path]; ok {
					dirsize += v
				}
				dirs[path] = dirsize
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	solution1 := uint64(0)
	for path := range dirs {
		dirsize := dirs[path]
		if dirsize <= 100000 {
			solution1 += dirsize
		}
	}
	fmt.Println(solution1)

	used := dirs[""]
	free := uint64(70000000) - used
	solution2 := uint64(0)
	for path := range dirs {
		dirsize := dirs[path]
		if free+dirsize >= 30000000 {
			if solution2 == 0 || dirsize < solution2 {
				solution2 = dirsize
			}
		}
	}
	fmt.Println(solution2)
}
