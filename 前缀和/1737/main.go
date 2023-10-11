package main

import "fmt"

func minCharacters(a string, b string) int {
	aArr := make([]int, 26)
	bArr := make([]int, 26)
	aSum := make([]int, 26)
	bSum := make([]int, 26)

	maxAcnt, maxBcnt := 0, 0

	for i := 0; i < len(a); i++ {
		idx := a[i] - 'a'
		aArr[idx]++
	}
	aSum[0] = aArr[0]
	for i := 1; i < len(aArr); i++ {
		aSum[i] = aSum[i-1] + aArr[i]
	}
	for i := 0; i < len(b); i++ {
		idx := b[i] - 'a'
		bArr[idx]++
	}
	bSum[0] = bArr[0]
	for i := 1; i < len(bArr); i++ {
		bSum[i] = bSum[i-1] + bArr[i]
	}
	ans := len(a) + len(b)
	for i := 1; i < 26; i++ {
		cur := 0
		cur += aSum[25] - aSum[i-1]
		cur += bSum[i-1]
		if cur < ans {
			ans = cur
		}
		cur = 0
		cur += bSum[25] - bSum[i-1]
		cur += aSum[i-1]
		if cur < ans {
			ans = cur
		}
	}

	for i := 0; i < 26; i++ {
		if bArr[i] > maxBcnt {
			maxBcnt = bArr[i]
		}
		if aArr[i] > maxAcnt {
			maxAcnt = aArr[i]
		}
	}

	cur := len(a) - maxAcnt + len(b) - maxBcnt
	if cur < ans {
		ans = cur
	}

	return ans
}

func main() {
	fmt.Println(minCharacters("otpaynexxlngyrdmand", "twtyifiiundfejxfktclktjnqrmycnqmxhxfitnlasyuwotjguerenjjnpkaajsnnraopdnambfccwthppimijghg"))
}
