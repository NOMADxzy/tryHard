package main

import "fmt"

func minimumBuckets(hamsters string) int { //TODO 极度易错
	if hamsters == "H" || len(hamsters) > 1 && (hamsters[:2] == "HH" || hamsters[len(hamsters)-2:] == "HH") {
		return -1
	}
	total := 0
	ans := 0
	already := 0
	if hamsters[0] == 'H' {
		total++
	}
	lastIdx := -1
	for i := 1; i < len(hamsters)-1; i++ {
		if hamsters[i] == 'H' {
			if hamsters[i-1] == 'H' && hamsters[i+1] == 'H' {
				return -1
			}
			total++
		} else {
			if lastIdx != i-1 && hamsters[i-1] == 'H' && hamsters[i+1] == 'H' {
				already += 2
				total += 1
				ans++
				lastIdx = i + 1
				i++
			}
		}
	}
	if lastIdx != len(hamsters)-1 && hamsters[len(hamsters)-1] == 'H' {
		total++
	}

	ans += total - already
	return ans
}

func main() {
	fmt.Println(minimumBuckets("HH........"))
}
