# 给n个木棍，一把锯子，问在保证得到指定长度木棍的数量不少于k的情况下，木棍的最大长度是多少

from typing import List
from math import floor

class Solution:
    def maxLength(self , nums: List[int], k: int) -> int:
        left, right = 1, max(nums)

        def check(cur_len):
            acc = 0
            i = 0
            while acc < k and i < len(nums):
                acc += floor(nums[i] / cur_len)
                i += 1
            return acc >= k

        if check(right):
            return right
        else:
            while left<right:
                mid = (left+right+1)//2
                if check(mid):
                    left = mid
                else:
                    right = mid-1
            return left


print(Solution().maxLength([1,2,3,4,5], 2))