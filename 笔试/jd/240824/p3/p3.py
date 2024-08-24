import heapq
from collections import deque


def is_valid(cnts: dict):
    tmp = []
    v1,v2,v3 = None,None,None
    for k,v in cnts.items():
        for _ in range(v):
            if len(tmp) < 3:
                tmp.append(k)
            elif v1==None:
                tmp.sort()
                v1,v2,v3 = tmp
            else:
                if k<v1:
                    v2 = v1
                    v1 = k
                elif k<v2:
                    v2 = k
                elif k>v3:
                    v3 = k
    if v1==None:
        tmp.sort()
        v1, v2, v3 = tmp
    return v1+v2 > v3



n = int(input())
lengths = list(map(int, input().split()))

# tmp = lengths[:3]
# tmp.sort()
# cur_min1, cur_min2, cur_max = tmp
l,r = 0,2
cnts = {}
for v in lengths[:2]:
    if v not in cnts:
        cnts[v] = 0
    cnts[v] += 1

max_len = 2
start = 0

while r<n:
    if lengths[r] not in cnts:
        cnts[lengths[r]] = 0
    cnts[lengths[r]] += 1
    while r-l+1>2 :
        if is_valid(cnts):
            break
        l += 1

    if r-l+1 > max_len:
        start = l
        max_len = r-l+1

    r+=1

print(start, start+max_len-1)

# 2 3 3 3 1 1 3 3 3