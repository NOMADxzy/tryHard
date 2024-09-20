
m, n, p, q = map(int, input().split())
# grid = [[0 for _ in range(n)] for _ in range(m)]

horizons = [[(i,0)] for i in range(m)]
verticals = [[(0,i)] for i in range(n)]
oblique0_top = [[(i, 0)] for i in range(m)]
oblique0_bottom = [[(0,i)] for i in range(n-1)]
oblique1_top = [[()] for i in range(m)]
oblique1_bottom = [[] for _ in range(n-1)]

for _ in range(p):
    i, j = map(int, input().split())
    # grid[i-1][j-1] = 1


for _ in range(q):
    u, v, t = map(int, input().split())

