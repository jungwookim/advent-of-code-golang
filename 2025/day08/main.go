package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
	z int
}

type MyData struct {
	p Position
	q Position
	d float64
}

type PositionCount struct {
	p Position
	c int
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

		if len(splitted_str) != 3 {
			return nil, errors.New("Input line size should be 3")
		}

		x, _ := strconv.Atoi(splitted_str[0])
		y, _ := strconv.Atoi(splitted_str[1])
		z, _ := strconv.Atoi(splitted_str[2])

		data = append(data, Position{x, y, z})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func prepare(data []Position) []MyData {
	var result []MyData
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			p := data[i]
			q := data[j]
			d := distance(p, q)
			result = append(result, MyData{p, q, d})
		}
	}
	return result
}

func distance(p Position, q Position) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2) + math.Pow(float64(p.z-q.z), 2))
}

var parent = make(map[Position]Position)

func find(p Position) Position {
	if p == parent[p] {
		return p
	}
	parent[p] = find(parent[p]) // optimization
	return parent[p]
}

func union(p Position, q Position) {
	x := find(p)
	y := find(q)

	parent[x] = y
}

func top3() (int, error) {
	temp := make(map[Position]int)

	for p, _ := range parent {
		temp[find(p)]++
	}

	var positionCount []PositionCount
	for p, c := range temp {
		positionCount = append(positionCount, PositionCount{p, c})
	}

	sort.Slice(positionCount, func(i, j int) bool {
		return positionCount[i].c > positionCount[j].c
	})

	if len(positionCount) < 3 {
		return -1, errors.New("Invalid answer")
	}

	ans := 1
	for _, pos := range positionCount[0:3] {
		ans *= pos.c
	}
	return ans, nil
}

func part1(maxConnection int) {
	parsed_data, _ := parse()
	for _, node := range parsed_data {
		parent[node] = node
	}
	data := prepare(parsed_data)

	sort.Slice(data, func(i, j int) bool {
		return data[i].d < data[j].d
	})

	cnt := 0
	for _, each := range data {
		if cnt >= maxConnection {
			break
		}
		p := each.p
		q := each.q
		if find(p) != find(q) {
			union(p, q)
		}
		cnt++
	}

	ans, err := top3()
	if err != nil {
		fmt.Println("There are no top3")
		return
	}
	fmt.Println("Answer:", ans)
	return

}

func main() {
	part1(10)
}
