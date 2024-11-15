# n的数组，每次选前x个大于0的数，统一减去一个min(arr[:x])，求重复若干次后求减去的数总和

n = int(input())
arr = list(map(int, input().split()))

dp_min_poses = [0] * n

for i in range(1, n):
    if arr[i] < arr[dp_min_poses[i-1]]:
        dp_min_poses[i] = i
    else:
        dp_min_poses[i] = dp_min_poses[i-1]

ans = 0
acc_del_val = 0
pos = n-1
# print(dp_min_poses)
while True:
    next_pos, del_val = dp_min_poses[pos], arr[dp_min_poses[pos]]
    ans += (pos+1) * (del_val - acc_del_val)
    acc_del_val = del_val
    if next_pos == 0:
        break
    pos = next_pos-1

print(ans)

