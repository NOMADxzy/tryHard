
arr1 = list(map(int, input().split()))
arr2 = list(map(int, input().split()))

# 1 3 5
# 2 4 6

arr3 = []
i, j = 0, 0
m, n = len(arr1), len(arr2)

while i<m and j<n:
    if arr1[i] <= arr2[j]:
        arr3.append(arr1[i])
        i += 1
    else:
        arr3.append(arr2[j])
        j += 1

arr3.extend(arr1[i:])
arr3.extend(arr2[j:])

size = m+n
ans = 0
if size % 2 == 0:
    ans = (arr3[size//2] + arr3[size//2-1])//2
else:
    ans = arr3[size//2]

print(ans)