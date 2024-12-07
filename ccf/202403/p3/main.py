
# 6
# 2 o2 o3
# 3 c1o1 c1o2 o2
# 2 n2o4 n1o2
# 4 cu1 h1n1o3 cu1n2o6 h2o1
# 4 al2s3o12 n1h5o1 al1o3h3 n2h8s1o4
# 4 c1o1 c1o2 o2 h2o1

n = int(input())
for _ in range(n):
    split_strs = input().split()
    m = int(split_strs[0]) # 变量数量
    split_strs = split_strs[1:]
    matrix_dict = {}
    for i in range(m):
        s = split_strs[i]
        j = 0
        while j<len(s):
            t = j
            while ord(s[t])>57:
                t += 1
            cur_meta = s[j:t]
            j = t
            while t<len(s) and ord(s[t])<=57:
                t += 1
            cur_cnt = int(s[j:t])
            if cur_meta not in matrix_dict:
                matrix_dict[cur_meta] = [0]*m
            matrix_dict[cur_meta][i] = cur_cnt
            j = t

    matrix = list(matrix_dict.values())
    n = len(matrix) # 元素数量/方程数量
    if n<m:
        print("Y")
        continue
    cur_row_idx = 0
    cur_col_idx = 0
    while cur_col_idx<m:
        if matrix[cur_row_idx][cur_col_idx] == 0:
            new_row_idx = cur_row_idx+1
            while new_row_idx<n and matrix[new_row_idx][cur_col_idx] == 0:
                new_row_idx += 1
            if new_row_idx==n:
                cur_col_idx += 1
                if cur_col_idx == m:
                    cur_row_idx -= 1
                    break
                continue
            else:
                matrix[cur_row_idx], matrix[new_row_idx] = matrix[new_row_idx], matrix[cur_row_idx]

        for i in range(cur_row_idx+1, n):
            times_num = matrix[i][cur_col_idx] / matrix[cur_row_idx][cur_col_idx]
            for j in range(cur_col_idx, m):
                matrix[i][j] -= matrix[cur_row_idx][j] * times_num
        cur_col_idx += 1
        if cur_col_idx == m:
            break
        cur_row_idx += 1

    zhi = cur_row_idx+1
    print("Y" if zhi<m else "N")
