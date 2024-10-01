from idlelib.tree import TreeNode

n, L, S = map(int, input().split())
trees = {}
points = []
for _ in range(n):
    x, y = map(int, input().split())
    points.append((x, y))
    if x not in trees:
        trees[x] = {}
    trees[x][y] = True

def is_tree(x, y):
    if x in trees and trees[x].get(y, False):
        return True
    else:
        return False

small_map = [None for _ in range(S+1)]
for i in range(S+1):
    small_map[S - i] = list(map(lambda x: bool(int(x)), input().split()))

ans = 0

for start_x, start_y in points:
    if start_x + S <= L and start_y + S <= L:
        ok = True
        for i in range(S+1):
            for j in range(S+1):
                a = small_map[i][j]
                b = is_tree(start_x+i, start_y+j)
                if a != b:
                    ok = False
                    break
        if ok:
            ans += 1

print(ans)