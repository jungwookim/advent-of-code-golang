package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parse() []map[string]int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var maps []map[string]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		direction := string(line[0])
		numStr := line[1:]
		value, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		maps = append(maps, map[string]int{direction: value})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return maps
}

func part1(maps []map[string]int) int {
	start := 50
	ans := 0

	for _, m := range maps {
		for dir, val := range m {
			if dir == "L" {
				start -= val
				if start < 0 {
					start = (start + 100) % 100
				}
			} else {
				// dir == "R"
				start = (start + val) % 100
			}
		}
		if start == 0 {
			ans++
		}
	}
	return ans
}

func part2(maps []map[string]int) int {
	start := 50
	ans := 0

	for _, m := range maps {
		for dir, val := range m {
			if dir == "L" {
				distanceToZero := start % 100
				if start == 0 {
					distanceToZero = 100
				}
				if distanceToZero <= val {
					hits := 1 + (val-distanceToZero)/100
					ans += hits
				}
				start = (start - val) % 100
				if start < 0 {
					start += 100
				}
			} else {
				distanceToZero := (100 - start) % 100
				if start == 0 {
					distanceToZero = 100
				}
				if distanceToZero <= val {
					hits := 1 + (val-distanceToZero)/100
					ans += hits
				}
				start = (start + val) % 100
			}
		}

	}
	return ans
}

func main() {
	maps := parse()

	res1 := part1(maps)
	fmt.Println(res1)

	res2 := part2(maps)
	fmt.Println(res2)
}
