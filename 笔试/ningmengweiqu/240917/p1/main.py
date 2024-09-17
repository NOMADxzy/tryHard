#
# 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
#
#
# @param nums int整型一维数组
# @return int整型一维数组
#
class Solution:
    def findNext(self , nums: List[int]) -> List[int]:
        # write code here
        ans = [-1 for _ in range(len(nums))]
        n = len(nums)
        right_arr = [nums[-1]]
        for i in range(n-2, -1, -1):
            cur = nums[i]
            if right_arr[-1] <= cur:
                right_arr.append(cur)
                ans[i] = -1
            else:
                l, r = 0, len(right_arr)
                while l<r:
                    mid = (l+r)//2
                    if right_arr[mid] > cur:
                        r = mid
                    else:
                        l = mid + 1
                ans[i] = right_arr[r]
                right_arr.insert(r, cur)
        return ans