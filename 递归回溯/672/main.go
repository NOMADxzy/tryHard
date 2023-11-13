package main

import "fmt"

type state struct {
	l1 int
	l2 int
	l3 int
	l4 int
}

func findChoosed(combines *[][]int, choosed []int, k int, lastPos int, limit int) {
	if k == 0 {
		tmp := make([]int, 4)
		copy(tmp, choosed)
		*combines = append(*combines, tmp)
		return
	}
	for i := lastPos + 1; i < limit; i++ {
		if choosed[i] == 0 {
			choosed[i] = 1
			findChoosed(combines, choosed, k-1, i, limit)
			choosed[i] = 0
		}
	}
}

func flipLights(n int, presses int) int {
	if presses == 0 {
		return 1
	}
	set := map[state]bool{}
	ans := 0
	choosed := make([]int, 4)

	for k := 0; k <= 4; k++ {
		if presses >= k && (presses-k)%2 == 0 {
			var combines [][]int

			findChoosed(&combines, choosed, k, -1, 4)
			for _, combine := range combines {
				st := state{combine[0] ^ combine[2] ^ combine[3], combine[0] ^ combine[1],
					combine[0] ^ combine[2], combine[0] ^ combine[1] ^ combine[3]}
				if n < 4 {
					st.l4 = 0
				}
				if n < 3 {
					st.l3 = 0
				}
				if n < 2 {
					st.l2 = 0
				}
				if !set[st] {
					set[st] = true
					ans++
				}

			}
		}
	}
	return ans

}

func main() {
	fmt.Println(flipLights(3, 2))
}
