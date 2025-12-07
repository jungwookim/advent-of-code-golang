package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines [][]string

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		lines = append(lines, parts)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func part1() int {
	lines := parse()
	operators := lines[len(lines)-1]
	results := make([]int, len(operators))
	for i, op := range operators {
		if op == "*" {
			results[i] = 1
		} else {
			results[i] = 0
		}
	}

	for _, line := range lines[0 : len(lines)-1] {
		for j, e := range line {
			n, _ := strconv.Atoi(e)
			if operators[j] == "*" {
				results[j] *= n
			} else {
				results[j] += n
			}
		}
	}

	ans := 0
	for _, res := range results {
		ans += res
	}
	return ans
}

func main() {
	fmt.Println(part1())
}
