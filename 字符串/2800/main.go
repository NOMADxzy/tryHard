package main

import "fmt"

func minimumString(a string, b string, c string) string {
	cpt := func(a, b string) (string, string) {
		ans1 := a + b
		ans2 := b + a
		for i := 0; i < len(a); i++ {
			delta := min(len(a)-i, len(b))
			if a[i:i+delta] == b[:delta] {
				if i+len(b) < len(a) {
					ans1 = a
				} else {
					ans1 = a[:i] + b
				}
				break
			}
		}
		for i := 0; i < len(b); i++ {
			delta := min(len(b)-i, len(a))
			if b[i:i+delta] == a[:delta] {
				if i+len(a) < len(b) {
					ans2 = b
				} else {
					ans2 = b[:i] + a
				}
				break
			}
		}
		return ans1, ans2
	}
	ans := a + b + c
	update := func(s string) {
		if len(s) < len(ans) || len(s) == len(ans) && s < ans {
			ans = s
		}
	}

	combs := [][]string{{a, b, c}, {b, c, a}, {a, c, b}}
	for _, comb := range combs {
		pre1, pre2 := cpt(comb[0], comb[1])
		tmpAns1, tmpAns2 := cpt(pre1, comb[2])
		update(tmpAns1)
		update(tmpAns2)
		tmpAns1, tmpAns2 = cpt(pre2, comb[2])
		update(tmpAns1)
		update(tmpAns2)
	}
	return ans
}

func main() {
	fmt.Println(minimumString("a", "a", "b"))
}
