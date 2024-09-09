
s = input()
k = int(input())

# HSDC
arr = []
for c in s:
    if c == 'H':
        arr.append(1)
    elif c == 'S':
        arr.append(2)
    elif c == 'D':
        arr.append(3)
    elif c == 'C':
        arr.append(4)

limit = len(s) - k
total = sum(arr)

l,r = 0, min(len(arr)-1, limit-1)
cur_sum = sum(arr[:r+1])
min_sum = cur_sum

r += 1

while r < len(s):
    cur_sum = cur_sum + arr[r] - arr[l]
    l += 1
    min_sum = min(min_sum, cur_sum)
    r += 1

print(total - min_sum)
