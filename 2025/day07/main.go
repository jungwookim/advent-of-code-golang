package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func parse() ([][]byte, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
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
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return grid, nil
}

func part1() int {
	grid, _ := parse()
	ans := 0

	row_size := len(grid)
	col_size := len(grid[0])
	for i := 1; i < row_size; i++ { // skip the first line
		for j := 0; j < col_size; j++ {
			if grid[i-1][j] == 'S' || grid[i-1][j] == '|' {
				if grid[i][j] == '^' {
					ans++
					grid[i][j-1] = '|'
					grid[i][j+1] = '|'
				} else {
					// case '.'
					grid[i][j] = '|'
				}
			}
		}
	}

	fmt.Println("Answer:", ans)

	return ans
}

func countTimelines(grid [][]byte, x int, y int, memo map[[2]int]int) int {
	// 좌우 경계를 벗어나면 유효하지 않은 경로
	if y < 0 || y >= len(grid[0]) {
		return 0
	}
	// 아래 경계를 벗어나면 완료
	if x >= len(grid) {
		return 1
	}

	// 메모이제이션: 이미 계산한 결과가 있으면 반환
	key := [2]int{x, y}
	if val, ok := memo[key]; ok {
		return val
	}

	var result int
	switch grid[x][y] {
	case 'S', '.':
		result = countTimelines(grid, x+1, y, memo)
	case '^':
		result = countTimelines(grid, x+1, y-1, memo) + countTimelines(grid, x+1, y+1, memo)
	default:
		result = 0
	}

	memo[key] = result
	return result
}

func findStart(grid [][]byte) (int, int, error) {
	for y, v := range grid[0] {
		if v == 'S' {
			return 0, y, nil
		}
	}
	return 0, 0, errors.New("start position 'S' not found")
}

func part2() int {
	grid, err := parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
		return -1
	}

	x, y, err := findStart(grid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return -1
	}
	memo := make(map[[2]int]int) // 메모 초기화

	ans := countTimelines(grid, x, y, memo)

	fmt.Println("Answer:", ans)
	return ans
}

func main() {
	part1()
	part2()
}

// What I learned:
// DFS
// memoization
// switch statement
// handling error
// these can be nil: pointer, slice, map, channel, function, interface
// these cannot be nil: int, float, bool, string, struct, array
