package main

import "fmt"

func pushDominoes(dominoes string) string {
	m := len(dominoes)
	dpToLeft := make([]int, m)
	dpToRight := make([]int, m)

	for i := 0; i < m; i++ {
		dpToRight[i] = -1
		if dominoes[i] == 'R' {
			dpToRight[i] = 0
		} else if dominoes[i] == '.' {
			if i > 0 && dpToRight[i-1] >= 0 {
				dpToRight[i] = dpToRight[i-1] + 1
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		dpToLeft[i] = -1
		if dominoes[i] == 'L' {
			dpToLeft[i] = 0
		} else if dominoes[i] == '.' {
			if i < m-1 && dpToLeft[i+1] >= 0 {
				dpToLeft[i] = dpToLeft[i+1] + 1
			}
		}
	}

	ans := make([]byte, m)

	for i := 0; i < m; i++ {
		if dominoes[i] == '.' {
			if dpToLeft[i] == -1 && dpToRight[i] == -1 {
				ans[i] = '.'
			} else if dpToLeft[i] == -1 {
				ans[i] = 'R'
			} else if dpToRight[i] == -1 {
				ans[i] = 'L'
			} else {
				if dpToRight[i] < dpToLeft[i] {
					ans[i] = 'R'
				} else if dpToRight[i] > dpToLeft[i] {
					ans[i] = 'L'
				} else {
					ans[i] = '.'
				}
			}
		} else {
			ans[i] = dominoes[i]
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(pushDominoes(".L.R...LR..L.."))
}
