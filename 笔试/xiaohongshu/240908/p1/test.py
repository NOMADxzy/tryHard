from functools import cmp_to_key

def better(a, b):
    if a<b:
        return -1
    elif a>b:
        return 1
    else:
        return 0

arr = [3,6,2,1,4,8,0]
arr.sort(key=cmp_to_key(better))
print(arr)