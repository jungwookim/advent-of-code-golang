package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var banks [][]int

	for scanner.Scan() {
		var bank []int
		line := scanner.Text()
		if line == "" {
			continue
		}

		for _, s := range line {
			bank = append(bank, int(s-'0'))
		}
		banks = append(banks, bank)
	}
	return banks
}

func part1(banks [][]int) {
	ans := 0
	for _, bank := range banks {
		lt := bank[0]
		rt := bank[1]
		for _, next := range bank[2:] {
			if lt < rt {
				lt = rt
				rt = next
			} else {
				if rt < next {
					rt = next
				}
			}
		}
		ans += (lt * 10) + rt
	}
	fmt.Println(ans)
}

func logic2(arr []int, k int) []int {
	n := len(arr)
	drop := n - k

	stack := []int{}

	for _, digit := range arr {
		for drop > 0 && len(stack) > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1]
			drop--
		}
		stack = append(stack, digit)
	}
	return stack[:k]
}

func digitsToInt(digits []int) int {
	var sb strings.Builder
	for _, d := range digits {
		sb.WriteByte(byte(d) + '0')
	}
	n, _ := strconv.Atoi(sb.String())
	return n
}

func part2(banks [][]int, k int) {
	ans := 0
	for _, bank := range banks {
		joltage := logic2(bank, k)
		ans += digitsToInt(joltage)
	}
	fmt.Println(ans)
}

func main() {
	banks := parse()

	part1(banks)
	// part1
	part2(banks, 2)
	part2(banks, 12)

	// i've got some problem of mutable code in part1. It affects banks that will be caculated in part2.
}
