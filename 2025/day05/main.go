package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func parse() ([]Range, []int) {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var idRanges []Range
	var ids []int
	for _, str := range strings.Fields(string(content)) {
		if strings.Contains(str, "-") {
			ss := strings.Split(str, "-")
			start, _ := strconv.Atoi(ss[0])
			end, _ := strconv.Atoi(ss[1])
			idRanges = append(idRanges, Range{Start: start, End: end})
		} else {
			id, _ := strconv.Atoi(str)
			ids = append(ids, id)
		}
	}

	return idRanges, ids
}

func part1() int {
	idRanges, ids := parse()

	ans := 0
	for _, id := range ids {
		for _, id_range := range idRanges {
			if id_range.Start <= id && id <= id_range.End {
				ans++
				break
			}
		}
	}
	return ans
}

func part2() int {
	idRanges, _ := parse()
	sort.Slice(idRanges, func(i, j int) bool {
		if idRanges[i].Start == idRanges[j].Start {
			return idRanges[i].End < idRanges[j].End
		}
		return idRanges[i].Start < idRanges[j].Start
	})
	curStart := idRanges[0].Start
	curEnd := idRanges[0].End
	var res []Range
	for i, idRange := range idRanges {
		if i == 0 {
			continue
		}
		nextStart := idRange.Start
		nextEnd := idRange.End
		if curStart <= nextStart && nextStart <= curEnd {
			curEnd = max(curEnd, nextEnd)
		} else {
			res = append(res, Range{Start: curStart, End: curEnd})
			curStart = nextStart
			curEnd = nextEnd
		}
	}
	res = append(res, Range{Start: curStart, End: curEnd})

	ans := 0
	for _, r := range res {
		ans += r.End - r.Start + 1
	}

	return ans
}

func main() {

	fmt.Println(part1())
	fmt.Println(part2())
}
