package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func parse() ([]Position, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []Position

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		splitted_str := strings.Split(line, ",")

		if len(splitted_str) != 2 {
			return nil, errors.New("Input line size should be 3")
		}

		x, _ := strconv.Atoi(splitted_str[0])
		y, _ := strconv.Atoi(splitted_str[1])

		data = append(data, Position{x, y})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func absInt(x int, y int) int {
	if x >= y {
		return x - y
	}
	return y - x
}

func area(p Position, q Position) int {
	dx := absInt(p.x, q.x) + 1
	dy := absInt(p.y, q.y) + 1
	return dx * dy
}

func part1() {
	tiles, _ := parse()

	var areas []int
	for i := 0; i < len(tiles); i++ {
		for j := 1; j < len(tiles); j++ {
			areas = append(areas, area(tiles[i], tiles[j]))
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(areas)))

	fmt.Println("Answer:", areas[0])
}

func main() {
	part1()
}
