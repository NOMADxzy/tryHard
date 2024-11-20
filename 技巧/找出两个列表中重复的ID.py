arr1 = [1, 3, 5, 7]
arr2 = [2, 3, 6, 7]

arr1.sort()
arr2.sort()

i, j = 0, 0
m, n = len(arr1), len(arr2)
ans = []

while i < m and j < n:
    if arr1[i] == arr2[j]:
        ans.append(arr1[i])
        i += 1
        j += 1
    elif arr1[i] < arr2[j]:
        i += 1
    else:
        j += 1

print(*ans)
