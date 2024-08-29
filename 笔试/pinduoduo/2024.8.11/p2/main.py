import heapq
import sys
from collections import deque
from math import inf

input()

n = int(input())

nums = []
for _ in range(n):
    nums.append(list(map(int, input().split())))

nums.sort(key=lambda x: (x[0], x[1]))
nums.append([inf, inf])
d = deque(nums)
# 结束时间，开始时间
h = []

res, now = 0,0

while d:
    start, remain = d.popleft()
    while h and now < start:
        h_r, h_s = heapq.heappop(h)
        if now + h_r <= start:
            now += h_r
            res += now - h_s
        else:
            res += start - h_s
            new_s, new_r = start, h_r - (start - now)
            heapq.heappush(h, (new_r, new_s))
            now = start
    now = start
    heapq.heappush(h, (remain, start))

print(res)