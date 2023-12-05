package main

import "fmt"

func findChildCnt(nexts [][]int, pos int, pre int, cntList []int) int {
	cnt := 1
	for _, p := range nexts[pos] {
		if p != pre {
			cnt += findChildCnt(nexts, p, pos, cntList)
		}
	}
	cntList[pos] = cnt
	return cnt
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads) + 1
	nexts := make([][]int, n)
	for _, road := range roads {
		nexts[road[0]] = append(nexts[road[0]], road[1])
		nexts[road[1]] = append(nexts[road[1]], road[0])
	}
	cntList := make([]int, n)
	findChildCnt(nexts, 0, 0, cntList)
	ans := int64(0)
	for i := 1; i < n; i++ {
		ans += int64(cntList[i] / seats)
		if cntList[i]%seats != 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumFuelCost([][]int{{1, 0}, {0, 2}, {3, 1}, {1, 4}, {5, 0}}, 1))
}
