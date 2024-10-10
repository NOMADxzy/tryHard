class Solution:
    def reviewAttendance(self, days: int) -> int:
        # 定义MOD常量，用来防止结果溢出
        MOD = 1000000007

        # dp[i][0] 表示以O结尾的合法序列数
        # dp[i][1] 表示以P结尾的合法序列数
        # dp[i][2] 表示以D结尾的合法序列数（前面没有D）
        # dp[i][3] 表示以DD结尾的合法序列数（即前一天是D，当前也是D）

        # 初始化dp数组
        dp = [[0] * 4 for _ in range(days)]
        dp[0] = [1, 1, 1, 0]  # 第一天可以是 O, P, D, 不能是 DD

        # 动态规划求解
        for i in range(1, days):
            dp[i][0] = (dp[i - 1][1] + dp[i - 1][2] + dp[i - 1][3]) % MOD  # 以O结尾，前一天不能是O，只能是P、D、DD
            dp[i][1] = (dp[i - 1][0] + dp[i - 1][1] + dp[i - 1][2] + dp[i - 1][3]) % MOD  # 以P结尾，前一天可以是任何状态
            dp[i][2] = (dp[i - 1][0] + dp[i - 1][1]) % MOD  # 以D结尾，前一天不能是D或DD，只能是O或P
            dp[i][3] = dp[i - 1][2]  # 以DD结尾，前一天必须是D

        # 计算最终结果，长度为 days 的序列可以以 O、P、D、DD 结尾
        return (dp[-1][0] + dp[-1][1] + dp[-1][2] + dp[-1][3]) % MOD


# 测试示例
solution = Solution()
days = 100
result = solution.reviewAttendance(days)
print(result)  # 输出序列的总数