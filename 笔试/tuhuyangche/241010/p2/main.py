from typing import List
# 背包
class Solution:
    def maxSpending(self , costs: List[int], budget: int) -> int:
        dp = {0}
        for x in costs:
            newDp = set(dp)
            for acc in dp:
                if acc+x <= budget:
                    newDp.add(acc+x)
            dp = newDp
        return max(dp)