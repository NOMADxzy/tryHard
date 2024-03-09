package main

import (
	"fmt"
)

func main() {
	var n, m, q int
	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&m)
	_, _ = fmt.Scan(&q)
	rMap := map[int]bool{}
	for i := 0; i < m; i++ {
		a, b := 0, 0
		_, _ = fmt.Scan(&a, &b)
		rMap[a*(n+1)+b] = true
	}
	for i := 0; i < q; i++ {
		tp, a, b := 0, 0, 0
		_, _ = fmt.Scan(&tp)
		_, _ = fmt.Scan(&a)
		_, _ = fmt.Scan(&b)
		if tp == 2 {
			if _, ok := rMap[a*(n+1)+b]; !ok {
				fmt.Println("No")
			} else {
				fmt.Println("Yes")
			}

		}
	}
}
