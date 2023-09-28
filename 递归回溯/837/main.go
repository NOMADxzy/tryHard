package main

import "fmt"

//内存溢出
func new21Game1(n int, k int, maxPts int) float64 {
	var states [][]int
	if 0 >= k {
		return 1
	}
	states = append(states, []int{0, 1})
	var probability float64

	for len(states) > 0 {
		st := states[0]
		states = states[1:]
		for pts := 1; pts <= maxPts; pts++ {
			if st[0]+pts < k {
				states = append(states, []int{st[0] + pts, st[1] * maxPts})
			} else if st[0]+pts <= n {
				probability += float64(1) / float64(st[1]*maxPts)
			}
		}
	}
	return probability
}

//递归法, 超时
func nextRound(maxPts int, acc, n, k int, visited map[int]float64) float64 {
	if visited[acc] > 0 {
		return visited[acc]
	}
	var res float64
	for pts := 1; pts <= maxPts; pts++ {
		if acc+pts < k {
			res += nextRound(maxPts, acc+pts, n, k, visited)
		} else if acc+pts <= n {
			res += 1
		} else {
			break
		}
	}
	visited[acc] = res / float64(maxPts)
	return res / float64(maxPts)
}

func new21Game2(n int, k int, maxPts int) float64 {
	if k == 0 {
		return 1
	}

	visited := map[int]float64{}
	return nextRound(maxPts, 0, n, k, visited)
}

//动态规划
func new21Game(n int, k int, maxPts int) float64 {
	maxScore := n
	maxWinScore := k - 1 + maxPts
	if k-1+maxPts > maxScore {
		maxScore = k - 1 + maxPts
		maxWinScore = n
	}
	dp := make([]float64, maxScore+1)
	for i := k; i <= maxWinScore; i++ {
		dp[i] = 1
	}
	for i := k - 1; i >= 0; i-- {
		sum := float64(0)
		for j := 1; j <= maxPts; j++ {
			sum += dp[i+j]
		}
		dp[i] = sum / float64(maxPts)
	}
	return dp[0]
}

//优化的动态规划
func new21Game_best(n int, k int, maxPts int) float64 {
	if k == 0 {
		return 1.0
	}
	dp := make([]float64, k+maxPts)
	for i := k; i <= n && i < k+maxPts; i++ {
		dp[i] = 1.0
	}

	dp[k-1] = 1.0 * float64(min(n-k+1, maxPts)) / float64(maxPts)
	for i := k - 2; i >= 0; i-- {
		dp[i] = dp[i+1] - (dp[i+maxPts+1]-dp[i+1])/float64(maxPts)
	}
	return dp[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(new21Game1(21, 17, 10))
}
