# 合并 K 个升序数组
# 给你一个二维数组，二维数组中的每个数组元素都已经按升序排列。

# 请你将所有数组合并到一个升序数组中，返回合并后的数组。

# 输入：lists = [[1,4,5],[1,3,4],[2,6]]
# 输出：[1,1,2,3,4,4,5,6]


def merge(arr_list):
    n = len(arr_list)
    if n == 1:
        return arr_list[0]
    left_list = arr_list[:n // 2]
    right_list = arr_list[n // 2:]

    left_res = merge(left_list)
    right_res = merge(right_list)
    len1, len2 = len(left_res), len(right_res)
    res = []
    i, j = 0, 0
    while i < len1 and j < len2:
        if left_res[i] <= right_res[j]:
            res.append(left_res[i])
            i += 1
        else:
            res.append(right_res[j])
            j += 1
    if i < len1:
        res.extend(left_res[i:])
    if j < len2:
        res.extend(right_res[j:])
    return res


ans = merge([[1, 4, 5], [1, 3, 4], [2, 6]])
print(ans)







