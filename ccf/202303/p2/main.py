import heapq
# 4 9 2
# 6 1
# 5 1
# 6 2
# 7 1

n, m, k = map(int, input().split())
detail_list = []

for _ in range(n):
    t, c = map(int, input().split())
    detail_list.append([-t, c])

heapq.heapify(detail_list)
while m > 0 and -detail_list[0][0] > k:
    elem = heapq.heappop(detail_list)

    if m<elem[1]:
        detail_list[0] = elem
        break
    m -= elem[1]
    elem[0] += 1
    heapq.heappush(detail_list, elem)

print(-detail_list[0][0])
