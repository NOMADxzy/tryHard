package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(sum int, k int, l, r int) (int, int) {
	minVal := sum + k*l
	maxVal := sum + k*r
	return minVal, maxVal
}

// 统计一下可能导致超时的原因：
// 1、没用long
// 2、读后面那q行时是先一个for循环存了，后面再来一个for循环输出的，而不是直接输出在读的时候就直接输出sum+l*count
// 3、输出时不是直接一个System.out.println(...+“ ”+...)而是分了多个
//
// 作者：牛客713799849号
// 链接：https://www.nowcoder.com/feed/main/detail/c3095b46f9e84f849303e78b0e6a39ef?sourceSSR=users
// 来源：牛客网
func main() {
	n, q := 0, 0
	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&q)
	var tmp int
	var sum int
	var k int
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	splits := strings.Split(scan.Text(), " ")
	for i := 0; i < n; i++ {
		tmp, _ = strconv.Atoi(splits[i])
		//_, _ = fmt.Scan(&tmp)
		if tmp == 0 {
			k++
		} else {
			sum += tmp
		}
	}
	ans := make([][2]int, q)
	for i := 0; i < q; i++ {
		l, r := 0, 0
		_, _ = fmt.Scan(&l)
		_, _ = fmt.Scan(&r)
		minVal, maxVal := solve(sum, k, l, r)
		ans[i][0] = minVal
		ans[i][1] = maxVal
	}
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i][0], ans[i][1])
	}
}
