
n = int(input())
arr = list(map(int, input().split()))
# 6
# 1 1 4 5 1 4

arr.sort()
new_arr_left, new_arr_right = [], []

i = 0
if len(arr)%2==1:
    new_arr_right.append(arr[0])
    i += 1

while i < len(arr):
    new_arr_left.append(arr[i])
    new_arr_right.append(arr[i+1])
    i += 2

new_arr_left.reverse()
new_arr = new_arr_left + new_arr_right

ans = 0
pre_total = 0
for i in range(n):
    cur_total = pre_total + 0 + new_arr[i]*(i+1)
    ans += cur_total
    pre_total = cur_total

# print(new_arr)
print(ans)