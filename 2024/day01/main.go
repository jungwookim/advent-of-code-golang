package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func intAbs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func main() {
	lt, rt := parse()
	part1(lt, rt)
	part2(lt, rt)
}

func part1(lt []int, rt []int) {
	sort.Ints(lt)
	sort.Ints(rt)

	var res int64 = 0
	for i := range len(lt) {
		res += int64(intAbs(lt[i] - rt[i]))
	}

	fmt.Println(res)
}

func part2(lt []int, rt []int) int64 {
	// counts := make(map[int]int) similar expressions, good when you want to set capacity
	counts := map[int]int{} // more idiomatic

	for i := 0; i < len(rt); i++ {
		counts[rt[i]]++
	}

	var ans int64

	for i := range len(lt) {
		ans += int64(lt[i] * counts[lt[i]])
	}

	fmt.Println(ans)

	return ans
}

func parse() ([]int, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lt []int
	var rt []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		l, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		lt = append(lt, l)
		rt = append(rt, r)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lt, rt
}
