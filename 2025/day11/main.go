package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parse(fileName string) (map[string][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	edges := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		splitted_str := strings.Split(line, ":")
		tos := strings.Fields(splitted_str[1])

		for _, to := range tos {
			edges[splitted_str[0]] = append(edges[splitted_str[0]], to)
		}

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return edges, nil
}

func process1(node string, visited map[string]bool, edges map[string][]string) int {
	if node == "out" {
		return 1
	}

	if visited[node] {
		return 0
	}

	visited[node] = true
	result := 0
	for _, next := range edges[node] {
		result += process1(next, visited, edges)
	}
	visited[node] = false
	return result
}

func process2(node string, visited map[string]bool, edges map[string][]string, depth int) int {
	if node == "out" {
		if visited["fft"] && visited["dac"] {
			return 1
		} else {
			return 0
		}
	}

	if visited[node] {
		return 0
	}

	visited[node] = true
	result := 0
	for _, next := range edges[node] {
		result += process2(next, visited, edges, depth+1)
	}
	visited[node] = false
	return result
}

func part1() {
	edges, _ := parse("input1.txt")
	visited := map[string]bool{}
	ans := process1("you", visited, edges)
	fmt.Println("Answer Part1:", ans)
}

func part2() {
	edges, _ := parse("input2.txt")
	visited := map[string]bool{}
	ans := process2("svr", visited, edges, 0)
	fmt.Println("Answer Part2:", ans)
}

func main() {
	part1()
	part2()
}
