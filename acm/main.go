package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func test5() {
	var n int
	for {
		_, _ = fmt.Scan(&n)
		if n == 0 {
			break
		}
		sum := 0
		var tmp int
		for i := 0; i < n; i++ {
			_, _ = fmt.Scan(&tmp)
			sum += tmp
		}
		fmt.Println(sum)
	}
}

func test7() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		splits := strings.Split(input.Text(), " ")
		sum := 0
		for i := 0; i < len(splits); i++ {
			if val, err := strconv.Atoi(splits[i]); err == nil {
				sum += val
			}
		}
		fmt.Println(sum)
	}
}

func main() {
	test7()
}
