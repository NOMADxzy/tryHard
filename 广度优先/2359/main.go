package main

import "fmt"

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	if node1 == node2 {
		return node1
	}
	n := len(edges)
	leftVisited, rightVisited := make([]bool, n), make([]bool, n)
	leftVisited[node1], rightVisited[node2] = true, true
	done1, done2, find := false, false, false
	ans := -1
	for (!done1 || !done2) && !find {
		if !done1 {
			if edges[node1] >= 0 && !leftVisited[edges[node1]] {
				node1 = edges[node1]
				leftVisited[node1] = true
				if rightVisited[node1] {
					find = true
					ans = node1
				}
			} else {
				done1 = true
			}
		}
		if !done2 {
			if edges[node2] >= 0 && !rightVisited[edges[node2]] {
				node2 = edges[node2]
				rightVisited[node2] = true
				if leftVisited[node2] {
					find = true
					if ans == -1 || ans > node2 {
						ans = node2
					}
				}
			} else {
				done2 = true
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(closestMeetingNode([]int{4, 4, 8, -1, 9, 8, 4, 4, 1, 1}, 5, 6))
}
