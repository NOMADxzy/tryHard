from typing import List
import heapq
# 一个会议列表，[[start,end]], 问最多可以参加多少个会议

class Solution:
    def attendmeetings(self , meetings: List[List[int]]) -> int:
        in_meetings = {}
        max_t = 0
        def add_point(a,b):
            if a not in in_meetings:
                in_meetings[a] = []
            in_meetings[a].append(b)

        for m in meetings:
            a, b = m
            add_point(a, b)
            max_t = max(max_t, b)

        h_arr = []
        ans = 0
        for cur in range(0, max_t+1):
            if cur in in_meetings:
                for e in in_meetings[cur]:
                    heapq.heappush(h_arr, e)
            while len(h_arr) > 0 and h_arr[0] < cur:
                heapq.heappop(h_arr)
            if len(h_arr)>0:
                heapq.heappop(h_arr)
                ans += 1
        return ans



print(Solution().attendmeetings([[1,2],[2,3],[3,4]]))
