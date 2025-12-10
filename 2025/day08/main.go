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

type Edge struct {
	p Position
	q Position
	d float64
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

func prepare(data []Position) []Edge {
	var result []Edge
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			p := data[i]
			q := data[j]
			d := distance(p, q)
			result = append(result, Edge{p, q, d})
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
	circuitSize := make(map[Position]int)

	for p, _ := range parent {
		circuitSize[find(p)]++
	}

	// make(type, length, capacity)
	sizes := make([]int, 0, len(circuitSize))
	for _, size := range circuitSize {
		sizes = append(sizes, size)
	}

	// classic
	// sort.Slice(sizes, func(i, j int) bool {
	// 	return sizes[i] > sizes[j]
	// })

	// use Sort method
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	if len(sizes) < 3 {
		return -1, errors.New("Invalid answer")
	}

	return sizes[0] * sizes[1] * sizes[2], nil
}

func part1() {
	positions, _ := parse()
	for _, node := range positions {
		parent[node] = node
	}
	edges := prepare(positions)

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].d < edges[j].d
	})

	maxConnection := 10
	if len(positions) >= 100 {
		maxConnection = 1000
	}
	cnt := 0
	for _, edge := range edges {
		if cnt >= maxConnection {
			break
		}
		p := edge.p
		q := edge.q
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

func oneLargeCircuit() bool {
	circuitSize := make(map[Position]int)

	for p, _ := range parent {
		circuitSize[find(p)]++
	}

	// make(type, length, capacity)
	sizes := make([]int, 0, len(circuitSize))
	for _, size := range circuitSize {
		sizes = append(sizes, size)
	}

	if len(sizes) == 1 {
		return true
	}
	return false
}

func part2() {
	positions, _ := parse()
	for _, node := range positions {
		parent[node] = node
	}
	edges := prepare(positions)

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].d < edges[j].d
	})

	for _, edge := range edges {
		p := edge.p
		q := edge.q
		if find(p) != find(q) {
			union(p, q)
		}
		if oneLargeCircuit() {
			fmt.Println("Answer:", p.x*q.x)
			break
		}
	}

	return
}

func main() {
	part1()
	part2() // sample input
}
