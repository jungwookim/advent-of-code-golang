package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Left  string
	Right string
}

func parse() []Range {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ranges []Range

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		id_ranges := strings.Split(line, ",")
		for _, id_range := range id_ranges {
			splitted_ids := strings.Split(id_range, "-")
			ranges = append(ranges, Range{Left: splitted_ids[0], Right: splitted_ids[1]})
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return ranges
}

func Pow10(exp int) int {
	result := 1
	for range exp {
		result *= 10
	}
	return result
}

func part1(ranges []Range) {
	ans := 0
	for _, r := range ranges {
		left_size := len(r.Left)
		left_int, _ := strconv.Atoi(r.Left)
		right_int, _ := strconv.Atoi(r.Right)
		lv := left_int / Pow10(((left_size + 1) / 2))
		rv := right_int / Pow10(((left_size + 1) / 2))

		for i := lv; i <= rv; i++ {
			candidate, _ := strconv.Atoi(strconv.Itoa(i) + strconv.Itoa(i))
			if left_int <= candidate && candidate <= right_int {
				ans += candidate
			}
		}
	}

	fmt.Println("part1 ans:", ans)
}

func main() {
	ranges := parse()

	part1(ranges)
}
