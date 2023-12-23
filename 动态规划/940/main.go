package main

import "fmt"

// 递归超时
func distinctSubseqII_(s string) int {
	var dfs func(pos, r int) int
	MOD := 1000000007
	hist := map[int]int{}
	dfs = func(pos int, r int) int { //每层选择一个字母
		key := r*2001 + pos
		if v, ok := hist[key]; ok {
			return v
		}
		if pos == len(s)-1 {
			return 1
		}
		visited := map[byte]bool{}
		cnt := 1 // 表示不选后面的
		for i := pos + 1; i < len(s); i++ {
			if !visited[s[i]] {
				visited[s[i]] = true
				cnt = (cnt + dfs(i, r+1)) % MOD
			}
		}
		hist[key] = cnt
		return cnt
	}

	return dfs(-1, 0) - 1
}

func distinctSubseqII(s string) int {
	n := len(s)
	MOD := 1000000007

	ans := 0
	lastK := make([]int, 26)

	dp := make([]int, n)
	for i := 0; i < 26; i++ {
		lastK[i] = -1
	}
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 26; j++ {
			if lastK[j] != -1 {
				dp[i] = (dp[i] + dp[lastK[j]]) % MOD
			}
		}
		lastK[s[i]-'a'] = i
	}
	for i := 0; i < 26; i++ {
		if lastK[i] != -1 {
			ans = (ans + dp[lastK[i]]) % MOD
		}
	}

	return ans
}

func main() {
	fmt.Println(distinctSubseqII("pqhneptornperyemqhsmrtxuwythhfzndppwtemeikikahpbbahhwprqqvuejbouwmiapxkocqaaujurbfknyxwlznirsxhaomzdvwedmbpciocktoxvuvjkaxrgqtzxbvunjoxacsnddwrgmjsyfyvpzgautteiyvkewjgrwtzhhxducgelhynkiglywjziieiifwpjmikvfzobygsayyjuifgeqkaectlcjsujxwakfxsifojkrywztrobzqiukdkcbvizwlpfsiursmgrmamjddxcotlhiabmtkvvnbgwafkbrnlrjyenbvappfielbqtcriohxoqokqfmldhvligbdyayycilglrsvkpsotuqompmnxbfhowwdjblnwajyorvoicqcqfvtarilzzjspogulawprwovllxjgqvhqwijmmkivoulaxcfvdrepsibualsobnsncbkwkytkjxfydecibpycyxnutxyklhblwvzluiesmzsqzlkqfhubihvalzdkjlxtoqrtstnsiieyrwtgiphiiplefjhzslgpcfxowfaxedxcfewsxhgtkdcvzorjdywlnmrbrnkpxqufanmjjbwxoyaspyticsffkngclyutqjiznrabgsiwcokfadkqmudjppzxjwovdzemxjjtypznznkhtrpezbdsfguhcoulcmloaftccstolcmskqtoebpicepofldukiernmhqjwhytkzlmcpdlpnqwfqbextejmurscikaolwqwyloyijfjyktdrkgybvmimikielcpaxtkbickygdcbgemiksczafhcadpgroprdsnovdhqpdtwpaxgjfzylsiflntpszatogalcbobrhdlezkppdgbtmaewcmyijgpprxazfsuphwikmhbzdagxxgwlkegzbhhogvfiuoiouvvmldmnpjyvgmkyqgclrmbfiuvaxsuwbpznuvwqblyjkdkpbclwnozadlzgxhgkt"))
}
