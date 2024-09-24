
m, n = map(int, input().split())
grid = []
for _ in range(m):
    grid.append(input())

# 4 4
# abcd
# bdgf
# czke
# dcbb

class Point:
    def __init__(self, x, y):
        self.x, self.y = x,y
def dfs(p1: Point, p2: Point) -> int:
    tmp_dict1 = {}
    tmp_dict2 = {}
    if p1.x-1>=0:
        c = grid[p1.x-1][p1.y]
        if c not in tmp_dict1:
            tmp_dict1[c] = []
        tmp_dict1[c].append(Point(p1.x-1, p1.y))
    if p1.y-1>=0:
        c = grid[p1.x][p1.y-1]
        if c not in tmp_dict1:
            tmp_dict1[c] = []
        tmp_dict1[c].append(Point(p1.x, p2.y-1))
    if p2.x+1<m:
        c = grid[p2.x+1][p2.y]
        if c not in tmp_dict2:
            tmp_dict2[c] = []
        tmp_dict2[c].append(Point(p2.x+1, p2.y))
    if p2.y+1<n:
        c = grid[p2.x][p2.y+1]
        if c not in tmp_dict2:
            tmp_dict2[c] = []
        tmp_dict2[c].append(Point(p2.x, p2.y+1))

    best = 0
    for c1, p_list1 in tmp_dict1.items():
        if c1 in tmp_dict2:
            p_list2 = tmp_dict2[c1]
            for p1 in p_list1:
                for p2 in p_list2:
                    best = max(best, dfs(p1, p2) + 2)

    return best

ans = 0
for i in range(m):
    for j in range(n):
        ans = max(ans, dfs(Point(i,j), Point(i,j)) + 1)
        if i+1<m:
            ans = max(ans, dfs(Point(i, j), Point(i+1, j)) + 2)
        if j+1<n:
            ans = max(ans, dfs(Point(i, j), Point(i, j+1)) + 2)

print(ans)
# dfs(Point(3,0), Point(3,0))