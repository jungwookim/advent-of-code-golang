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

// 메모이제이션: "노드,fft,dac" -> 경로 수
var memo map[string]int

func makeKey(node string, fftVisited, dacVisited bool) string {
	return fmt.Sprintf("%s,%v,%v", node, fftVisited, dacVisited)
}

func process2(node string, fftVisited, dacVisited bool, edges map[string][]string) int {
	// 현재 노드가 fft 또는 dac이면 플래그 업데이트
	if node == "fft" {
		fftVisited = true
	}
	if node == "dac" {
		dacVisited = true
	}

	if node == "out" {
		if fftVisited && dacVisited {
			return 1
		}
		return 0
	}

	// 이미 계산한 상태라면 캐시된 값 반환
	key := makeKey(node, fftVisited, dacVisited)
	if val, ok := memo[key]; ok {
		return val
	}

	result := 0
	for _, next := range edges[node] {
		result += process2(next, fftVisited, dacVisited, edges)
	}

	// 결과 캐싱
	memo[key] = result

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
	// 메모이제이션 초기화
	memo = make(map[string]int)
	ans := process2("svr", false, false, edges)
	fmt.Println("Answer Part2:", ans)
}

func main() {
	part1()
	part2()
}
