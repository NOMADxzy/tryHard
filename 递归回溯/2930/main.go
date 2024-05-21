package main

import "fmt"

const mod = 1_000_000_007

func stringCount(n int) int {
	hist := map[int]int{}
	MOD := 1000000007
	var dfs func(lcnt, ecnt, tcnt int, limit int) int
	dfs = func(lcnt, ecnt, tcnt int, limit int) int {
		if limit == 0 {
			if lcnt == 0 && ecnt == 0 && tcnt == 0 {
				return 1
			}
		}
		if lcnt+ecnt+tcnt > limit {
			return 0
		}
		key := limit*12 + ecnt*4 + lcnt*2 + tcnt
		if v, ok := hist[key]; ok {
			return v
		}
		cnt := 0
		for i := 'a'; i <= 'z'; i++ {
			newLcnt, newEcnt, newTcnt := lcnt, ecnt, tcnt
			if i == 'l' && newLcnt > 0 {
				newLcnt--
			} else if i == 'e' && newEcnt > 0 {
				newEcnt--
			} else if i == 't' && newTcnt > 0 {
				newTcnt--
			}
			cnt = (cnt + dfs(newLcnt, newEcnt, newTcnt, limit-1)) % MOD
		}
		hist[key] = cnt
		return cnt
	}
	return dfs(1, 2, 1, n)
}

func stringCount1(n int) (ans int) {
	return ((pow(26, n)-
		pow(25, n-1)*(75+n)+
		pow(24, n-1)*(72+n*2)-
		pow(23, n-1)*(23+n))%mod + mod) % mod // 保证结果非负
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func stringCount_(n int) int {
	MOD := 1000000007
	ans := 0
	getVal := func(x int) int {
		r := 1
		for x > 0 {
			r = ((x % MOD) * (r % MOD)) % MOD
			x--
		}
		return r
	}
	mother := map[int]int{}
	child := map[int]int{}
	var meanVal, plusNum int
	analy := func(x int, cnt map[int]int) {
		for i := 2; i <= x; {
			if x%i == 0 {
				x /= i
				cnt[i]++
			} else {
				i++
			}
		}
	}
	if n < 4 {
		return 0
	} else if n-4 <= 23 {
		ans = getVal(n) / getVal(2)
		tmp := min(n-4, 23-(n-4))
		for i := 0; i < tmp; i++ {
			v := 23 - i
			ans = (ans * v) % MOD
		}
		return ans
	} else {
		meanVal = n / 26
		plusNum = n - meanVal*26

		for i := 2; i <= n; i++ {
			analy(i, child)
		}
		for j := 0; j < 26; j++ {
			for i := 2; i <= meanVal; i++ {
				analy(i, mother)
			}
		}

		for j := 0; j < plusNum; j++ {
			analy(meanVal+1, mother)
		}
		for v, c := range mother {
			child[v] -= c
		}
		for v, cnt := range child {
			for i := 0; i < cnt; i++ {
				ans = (ans * v) % MOD
			}
		}
		return ans
	}
}

func main() {
	fmt.Println(stringCount(10))
}
