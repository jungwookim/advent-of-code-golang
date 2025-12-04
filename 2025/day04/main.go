package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse() [][]byte {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		grid = append(grid, []byte(line))
	}
	return grid
}

func logic1(grid [][]byte, x int, y int) int {
	res := 0
	max_y := len(grid)
	max_x := len(grid[0])
	if y-1 >= 0 && y-1 < max_y && x-1 >= 0 && x-1 < max_x && grid[y-1][x-1] == '@' {
		res += 1
	}
	if y-1 >= 0 && y-1 < max_y && x >= 0 && x < max_x && grid[y-1][x] == '@' {
		res += 1
	}
	if y-1 >= 0 && y-1 < max_y && x+1 >= 0 && x+1 < max_x && grid[y-1][x+1] == '@' {
		res += 1
	}

	if y >= 0 && y < max_y && x-1 >= 0 && x-1 < max_x && grid[y][x-1] == '@' {
		res += 1
	}
	if y >= 0 && y < max_y && x+1 >= 0 && x+1 < max_x && grid[y][x+1] == '@' {
		res += 1
	}

	if y+1 >= 0 && y+1 < max_y && x-1 >= 0 && x-1 < max_x && grid[y+1][x-1] == '@' {
		res += 1
	}
	if y+1 >= 0 && y+1 < max_y && x >= 0 && x < max_x && grid[y+1][x] == '@' {
		res += 1
	}
	if y+1 >= 0 && y+1 < max_y && x+1 >= 0 && x+1 < max_x && grid[y+1][x+1] == '@' {
		res += 1
	}
	if res < 4 {
		return 1
	}
	return 0
}

func part1() int {
	grid := parse()
	ans := 0

	for i := range len(grid) {
		for j := range len(grid[i]) {
			if grid[i][j] == '@' {
				ans += logic1(grid, j, i)
			}
		}
	}

	return ans
}

func part2() int {
	grid := parse()
	ans := 0

	for {
		changed := false
		for i := range len(grid) {
			for j := range len(grid[i]) {
				if grid[i][j] == '@' {
					add := logic1(grid, j, i)
					if add == 1 {
						changed = true
						ans++
						grid[i][j] = '.'
					}
				}
			}
		}
		if !changed {
			break
		}
	}

	return ans
}

func main() {
	res1 := part1()
	fmt.Println(res1)

	res2 := part2()
	fmt.Println(res2)
}

// What I leanred:
// String is immutable
// So if I would like to change string value, I could use []byte.
