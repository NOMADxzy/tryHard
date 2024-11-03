from typing import List

arr = list(map(int, input().split()))


# 5 2 8 3 9 1 7

def quickSort(arr: List, start: int, end: int):
    if end - start <= 0:
        return
    pivot = arr[start]
    l, r = start, end
    while l < r:
        while l < r and arr[r] >= pivot:
            r -= 1
        arr[l], arr[r] = arr[r], arr[l]
        while l < r and arr[l] <= pivot:
            l += 1
        arr[r], arr[l] = arr[l], arr[r]

    quickSort(arr, start, r - 1)
    quickSort(arr, r + 1, end)


quickSort(arr, 0, len(arr) - 1)
print(arr)
