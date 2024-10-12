from typing import List
# 每项会议从start持续到end，问最多可以参加多少
class Solution:
    def attendmeetings(self , meetings: List[List[int]]) -> int:
        meetings.sort(key=lambda x:x[1])
        n = len(meetings)
        dp = [0] * n
        dp[0] = 1
        for i in range(1, n):
            left, right = 0, i+1
            target = meetings[i][0]
            while left<right:
                mid = (left+right+1)//2
                if meetings[mid-1][1] < target:
                    left = mid
                else:
                    right = mid-1
            v = 1 + (dp[left-1] if left>0 else 0)
            dp[i] = max(dp[i-1], v)

        return dp[n-1]

print(Solution().attendmeetings([[1,2],[2,3],[]]))