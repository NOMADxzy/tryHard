

n = int(input())

for _ in range(n):
    full_path, left, right = input().split()
    full_elems = full_path.split('/')
    left_elems = left.split('/')
    right_elems = right.split('/')

    i = 0
    j = 0
    while i<len(full_elems) and j<len(left_elems) and full_elems[i] == left_elems[j]:
        i += 1
        j += 1

    if j<len(left_elems):
        print(0)
        continue

    j = 0
    while i<len(full_elems) and j<len(right_elems) and full_elems[i] == right_elems[j]:
        i += 1
        j += 1

    if j<len(right_elems) or i<len(full_elems):
        print(0)
        continue

    print(1)