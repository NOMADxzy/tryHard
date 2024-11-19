# 统计出现次数超过k的字符个数
#
class Solution:
    def findKthTrace(self, mystring: str, k: int) -> int:
        # write code here
        cnts = {}
        ans = 0
        for c in mystring:
            if c not in cnts:
                cnts[c] = 0
            cnts[c] += 1

        for c, cnt in cnts.items():
            if cnt >= k:
                ans += 1
        return ans
