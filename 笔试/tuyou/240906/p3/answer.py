from sortedcontainers import SortedList

n, k = map(int, input().split())
arr = list(map(int, input().split()))

# 使用 SortedList 维护有序的窗口
window = SortedList(arr[:k])


# 计算窗口的中位数
def get_mid(window):
    if len(window) % 2 == 0:
        return (window[len(window) // 2 - 1] + window[len(window) // 2]) // 2
    else:
        return window[len(window) // 2]


ans = [get_mid(window)]
for i in range(k, len(arr)):
    new_val = arr[i]
    del_val = arr[i - k]

    # 插入新值，删除旧值
    window.add(new_val)
    window.remove(del_val)

    # 计算当前窗口的中位数
    ans.append(get_mid(window))

print(" ".join(map(str, ans)))