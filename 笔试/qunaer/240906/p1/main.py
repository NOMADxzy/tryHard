from functools import cmp_to_key

n = int(input())
arr = list(map(str, input().split()))


def better(a: str, b: str) -> int:
    i, j = 0, 0
    while i < len(a) and j < len(b):
        if a[i] < b[j]:
            return 1
        elif a[i] > b[j]:
            return -1
        i += 1
        j += 1

    if i != len(a):
        j = 0
        while i < len(a):
            if a[i] < a[j]:
                return 1
            elif a[i] > a[j]:
                return -1
            i += 1
            j += 1
    if j != len(b):
        i = 0
        while j < len(b):
            if b[i] < b[j]:
                return 1
            elif b[i] > b[j]:
                return -1
            i += 1
            j += 1
    return 0


def quick_sort(arr):
    if len(arr) <= 1:
        return arr
    pivot = arr[len(arr) // 2]
    left = [x for x in arr if better(x, pivot) > 0]
    middle = [x for x in arr if better(x, pivot) == 0]
    right = [x for x in arr if better(x, pivot) < 0]
    return quick_sort(left) + middle + quick_sort(right)


arr = quick_sort(arr)

print(" ".join(arr))


