# risk列表重复执行：1️⃣求平均数、2️⃣删除平均数以上的数
# 问最多可以执行多少次？
#
class Solution:
    def deleteRisks(self, n: int, risks: List[int]) -> int:
        # write code here
        risks.sort()
        cur_total = sum(risks)
        right = n - 1

        ans = 0
        while right > 0:
            cur_avg = cur_total / (right + 1)
            if risks[right] == cur_avg:
                break

            new_right = right - 1
            cur_total -= risks[right]
            while new_right >= 0 and risks[new_right] > cur_avg:
                cur_total -= risks[new_right]
                new_right -= 1
            ans += 1
            right = new_right
        return ans

