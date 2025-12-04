package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	for i := 0; i < len(banks); i++ {
		bank := banks[i]
		joltage := bank[0:2]

		for j := 0; j < len(bank[2:]); j++ {
			next := bank[j+2]
			lt := joltage[0]
			rt := joltage[1]
			if lt < rt {
				joltage[0] = joltage[1]
				joltage[1] = next
			} else {
				if rt < next {
					joltage[1] = next
				}
			}
		}
		fmt.Println("part1:", joltage)
		ans += (joltage[0] * 10) + joltage[1]
	}
	fmt.Println(ans)
}

func logic2(arr []int, k int) []int {
	fmt.Println("arr:", arr)
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
	fmt.Println("stack:", stack)
	return stack[:k]
}

func part2(banks [][]int, k int) {
	ans := 0
	for i := 0; i < len(banks); i++ {
		bank := banks[i]
		joltage := logic2(bank, k)
		fmt.Println("part2:", joltage)

		temp_str := ""

		for _, digit := range joltage {
			temp_str += strconv.Itoa(digit)
		}

		real_joltage, _ := strconv.Atoi(temp_str)

		ans += real_joltage
	}
	fmt.Println(ans)
}

func main() {
	banks := parse()

	part1(banks)
	// part1
	// part2(banks, 2)
	part2(banks, 12)

	// i've got some problem of mutable code in part1. It affects banks that will be caculated in part2.
}
