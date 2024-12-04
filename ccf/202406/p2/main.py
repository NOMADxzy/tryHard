from symbol import return_stmt


m, n, t = map(int, input().split())

# i,j -> (i*n+j)//q, (i*n+j)%q
arr = []
for _ in range(m):
    arr.append(list(map(int, input().split())))

flatten_arr = []
for row in arr:
    flatten_arr += row

row_main = True
d = n
total = m*n

for _ in range(t):
    opt, v1, v2 = map(int, input().split())
    if opt == 1:
        if not row_main:
            new_flatten_arr = []
            for i in range(d):
                for j in range(total//d):
                    new_flatten_arr.append(flatten_arr[j*d+i])
            row_main = True
        d = v2
        flatten_arr = new_flatten_arr
    elif opt == 2:
        row_main = not row_main

    else:
        i, j = v1, v2
        if row_main:
            print(flatten_arr[i*d+j])
        else:
            print(flatten_arr[j*d + i])