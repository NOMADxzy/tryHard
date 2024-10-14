
n = int(input())
b_arr = list(map(int, input().split()))

# 6
# 0 0 5 5 10 10
# 0 0 5 5 10 10
# 0 0 5 0 10 0

cur_1 = b_arr[0]
cur_2 = b_arr[0]

sum_1, sum_2 = cur_1, cur_2

for i in range(1, n):
    if b_arr[i] > b_arr[i-1]:
        cur_1 = b_arr[i]
        cur_2 = b_arr[i]
    else:
        cur_1 = 0

    sum_1 += cur_1
    sum_2 += cur_2

print(sum_2)
print(sum_1)
