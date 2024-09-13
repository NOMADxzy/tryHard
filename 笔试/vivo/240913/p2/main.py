from typing import List
# 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
#
#
# @param memoryUsage int整型一维数组 数组 memoryUsage，其中 memoryUsage[i] 表示应用在第 i 秒的内存使用量（以MB为单位）
# @param k int整型 整数k
# @return int整型一维数组
#
class Solution:
    def findFluctuations(self, memoryUsage: List[int], k: int) -> List[int]:
        ans = []
        l, r = 0, k - 1
        while r < len(memoryUsage):
            min_val, max_val = memoryUsage[l], memoryUsage[l]
            for j in range(l, r + 1):
                min_val = min(min_val, memoryUsage[j])
                max_val = max(max_val, memoryUsage[j])
            ans.append(max_val - min_val)
            l += 1
            r += 1

        return ans