n, max_t = map(int, input().split())
arr = []
for _ in range(n):
    arr.append(list(map(int, input().split())))

left, right = 0, max([max(row) for row in arr])
dirs = [[-1, 0], [1, 0], [0, -1], [0, 1]]


def check(cur):
    queue = [(0, 0)]
    visited = [[False] * n for _ in range(n)]
    visited[0][0] = True
    t = 0
    while len(queue) > 0:
        if t>max_t:
            return False
        next_q = []
        while len(queue) > 0:
            e = queue[0]
            if e[0] == n - 1 and e[1] == n - 1:
                return True
            queue = queue[1:]
            for dir in dirs:
                nx, ny = e[0] + dir[0], e[1] + dir[1]
                if 0 <= nx < n and 0 <= ny < n and (arr[nx][ny] == -1 or arr[nx][ny]>=cur) and not visited[nx][ny]:
                    visited[nx][ny] = True
                    next_q.append((nx, ny))

        queue = next_q
        t += 1


if 2*n-1 > max_t or not check(0):
    print(0)
elif check(right):
    print(right)
else:
    while left < right:
        mid = (left + right + 1) // 2

        if check(mid):
            left = mid
        else:
            right = mid-1

    print(left)

# 4 6
#-1 2 1 1
# 0 2 2 0
# 0 0 2 2
# 1 1 0 -1