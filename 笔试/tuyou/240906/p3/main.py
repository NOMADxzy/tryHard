import heapq
from bisect import bisect

n, k = map(int, input().split())
arr = list(map(int, input().split()))

# 8 3
# 2 2 2 3 3 3 3 4

window = arr[:k]
window.sort()


def get_mid(window):
    if len(window) % 2 == 0:
        return (window[len(window)//2-1] + window[len(window)//2]) // 2
    else:
        return window[len(window)//2]

ans = [get_mid(window)]
for i in range(k, len(arr)):
    new_val = arr[i]
    if new_val<=window[0]:
        window.insert(0, new_val)
    elif new_val>= window[-1]:
        window.append(new_val)
    else:
        l, r = 0, len(window)-1
        while l<r:
            mid = (l+r)//2
            if window[mid] >= new_val:
                r = mid
            else:
                l = mid + 1
        window.insert(r, new_val)
    del_val = arr[i-k]
    if del_val==window[0]:
        window = window[1:]
    elif del_val== window[-1]:
        window = window[:-1]
    else:
        l, r = 0, len(window)-1
        while l<r:
            mid = (l+r)//2
            if window[mid] == del_val:
                l = mid
                r = mid
            elif window[mid] < del_val:
                l = mid + 1
            elif window[mid] > del_val:
                r = mid - 1
        window = window[:r] + window[r+1:]

    # window.sort()
    ans.append(get_mid(window))
    assert len(window) == k

print(" ".join(map(str, ans)))
