package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse1() [][]string {
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
	lines := parse1()
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

func parse2() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func skip(cols []string) bool {
	for _, col := range cols {
		if col != " " {
			return false
		}
	}
	return true
}

func colsToInt(cols []string) (int, string) {
	var sb strings.Builder
	for _, col := range cols {
		if col == "+" {
			n, _ := strconv.Atoi(sb.String())
			return n, "+"
		}
		if col == "*" {
			n, _ := strconv.Atoi(sb.String())
			return n, "*"
		}
		if col != " " {
			sb.WriteString(col)
		}
	}
	n, _ := strconv.Atoi(sb.String())
	return n, ""
}

func mul(numbers []int) int {
	result := 1
	for i := 0; i < len(numbers); i++ {
		result *= numbers[i]
	}
	return result
}

func sum(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func part2() int {
	lines := parse2()
	ans := 0
	var ps []int

	for i := len(lines[0]) - 1; i >= 0; i-- {
		var cols []string
		for j := 0; j < len(lines); j++ {
			cols = append(cols, string(lines[j][i]))
		}
		if skip(cols) {
			ps = ps[:0]
		} else {
			n, op := colsToInt(cols)
			if op == "" {
				ps = append(ps, n)
			} else {
				ps = append(ps, n)
				if op == "*" {
					ans += mul(ps)
				} else {
					ans += sum(ps)
				}
			}

		}
	}

	return ans
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
