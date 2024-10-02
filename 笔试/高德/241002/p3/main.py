from typing import List
#有n个气球，每个气球水平方向的起点和终点为(sx,ex), 求最少射出多少支垂直方向的箭可以射爆所有气球
#
class Solution:
    def findMinArrowShots(self , points: List[List[int]]) -> int:
        points.sort()
        n = len(points)
        ans = 0
        i = 0
        while i < n:
            cur_right = float('inf')
            while i<n and points[i][0] <= cur_right:
                cur_right = min(cur_right, points[i][1])
                i += 1
            ans += 1

        return ans

# s = Solution()
# print(s.findMinArrowShots([[0,4],[1,2],[2,3],[3,4]]))