# 排列5个数形成w形式

arr = list(map(int, input().split()))
arr.sort()
ans = [arr[4],arr[0],arr[2],arr[1],arr[3]]
print(" ".join(map(str, ans)))