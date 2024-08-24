import sys


# 0 0 0 0 0
# 0 0 0 1 0
# 0 0 0 0 0
# 0 1 0 0 0
# 0 0 0 0 0


param = list(map(int, input().split()))
n,m,k = param

arr = [[False for _ in range(m)] for _ in range(n)]

for _ in range(k):
    l = input().split()
    opt, i, j = l[0], int(l[1]), int(l[2])
    i-=1
    j-=1
    if opt == "c":
        arr[i][j] = True
    elif opt == 'l':
        j-=1
        while j>=0 and arr[i][j]:
            j-=1
    elif opt == 'r':
        j+=1
        while j<m and arr[i][j]:
            j+=1
    elif opt == 'u':
        i-=1
        while i>=0 and arr[i][j]:
            i-=1
    elif opt == 'd':
        i+=1
        while i<n and arr[i][j]:
            i+=1
    if opt == "c":
        continue
    if i<0 or i>=n or j<0 or j>=m:
        print(-1)
    else:
        print(i+1, j+1)


# for line in sys.stdin:
#     a = line.split()
#     print(int(a[0]) + int(a[1]))
