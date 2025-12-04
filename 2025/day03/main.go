package main

import (
	"bufio"
	"fmt"
	"os"
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
		ans += (joltage[0] * 10) + joltage[1]
	}
	fmt.Println(ans)
}

func main() {
	banks := parse()

	part1(banks)
}
