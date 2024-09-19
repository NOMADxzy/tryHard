from typing import List


def quick_sort(arr: List, l: int=0, r: int=-1):
    if r == -1:
        r = len(arr)-1
    if r-l+1 <= 1:
        return
    pivot = arr[l]
    i, j = l, r
    while i<j:
        while i<j and arr[j]>=pivot:
            j -= 1
        arr[i], arr[j] = arr[j],arr[i]
        while i<j and arr[i]<=pivot:
            i += 1
        arr[i], arr[j] = arr[j], arr[i]

    quick_sort(arr, l, i-1)
    quick_sort(arr, i+1, r)


arr = [1,5,8,2,5,8,3,4,0,1]
quick_sort(arr)
print(arr)