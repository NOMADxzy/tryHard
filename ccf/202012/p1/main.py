
n = int(input())

sum_val = 0
for _ in range(n):
    w, s = map(int, input().split())
    sum_val += w * s

sum_val = max(0, sum_val)

print(sum_val)