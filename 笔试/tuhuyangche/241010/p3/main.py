from typing import List
# P、O、D构成长度为days的序列不能连续出现两个O，不能连续出现3个D，求有多少种情况
class Solution:
    def reviewAttendance(self , days: int) -> int:
        dp = [[0] * 4 for _ in range(days)]
        dp[0] = [1,1,1,0]
        # O
        # P
        # D
        # DD
        MOD = 1000000007
        for i in range(1, days):
            dp[i][0] = ((dp[i-1][1] + dp[i-1][2]) % MOD + dp[i-1][3]) % MOD
            dp[i][1] = ((dp[i-1][0] + dp[i-1][1]) % MOD + (dp[i-1][2] + dp[i-1][3]) % MOD) % MOD
            dp[i][2] = (dp[i-1][0] + dp[i-1][1]) % MOD
            dp[i][3] = dp[i-1][2]

        return ((dp[-1][0] + dp[-1][1]) % MOD + (dp[-1][2] + dp[-1][3]) % MOD) % MOD

print(Solution().reviewAttendance(3))

# DDD
# OO?
# ?OO
# DDD OOP OOD OOO POO DOO
#
# PO DO
#
# PP DP OP
#
# OD PD
#
# DD
#
# PPO DPO OPO ODO PDO DDO
# POP DOP PPP DPP OPP ODP PDP DDP
# POD DOD PPD DPD OPD
# ODD PDD
