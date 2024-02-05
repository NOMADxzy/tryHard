package main

import "fmt"

func smallestNumber(num int64) int64 {
	if num == 0 {
		return 0
	}

	cnt := make([]int, 10)
	k := 0
	i := num
	if i < 0 {
		i = -i
	}
	for ; i != 0; i /= 10 {
		cnt[i%10]++
		k++
	}
	ans := int64(0)
	for k > 0 {
		for j := 0; j < 10; j++ {
			v := j
			if num < 0 {
				v = 9 - j
			}
			if ans == 0 && v == 0 {
				continue
			} else if cnt[v] > 0 {
				cnt[v]--
				k--
				ans = 10*ans + int64(v)
				break
			}
		}
	}
	if num < 0 {
		ans = -ans
	}
	return ans
}

func main() {
	fmt.Println(smallestNumber(7605800999))
}
