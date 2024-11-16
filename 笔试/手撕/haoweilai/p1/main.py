from typing import List
from functools import cmp_to_key


class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        n = len(numbers)
        idxs = [i for i in range(n)]

        def better(a, b):
            if numbers[a] < numbers[b]:
                return -1
            elif numbers[a] > numbers[b]:
                return 1
            else:
                return 0

        idxs.sort(key=cmp_to_key(better))
        cur_sum = numbers[idxs[0]] + numbers[idxs[-1]]
        l, r = 0, n - 1
        while l < r:
            if cur_sum < target:
                cur_sum -= numbers[idxs[l]]
                l += 1
                cur_sum += numbers[idxs[l]]
            elif cur_sum > target:
                cur_sum -= numbers[idxs[r]]
                r -= 1
                cur_sum += numbers[idxs[r]]
            else:
                break
        ans = [idxs[l] + 1, idxs[r] + 1]
        if ans[0] > ans[1]:
            ans[0], ans[1] = ans[1], ans[0]
        return ans
