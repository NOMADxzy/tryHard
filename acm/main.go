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

func test8() {
	// 创建一个 Scanner 来读取标准输入
	scanner := bufio.NewScanner(os.Stdin)

	// 设置缓冲区和最大缓冲区大小
	const maxCapacity = 500000 // 例如，我们将最大容量设置为500KB
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	// 在这里使用 scanner 读取数据
	fmt.Println("Please enter a very long line of text:")
	if scanner.Scan() {
		input := scanner.Text()
		fmt.Printf("You entered: %s\n", input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from input:", err)
	}

}

func main() {
	test7()
}
