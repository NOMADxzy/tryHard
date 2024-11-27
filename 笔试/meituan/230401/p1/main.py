
n = int(input())
arr = list(map(int, input().split()))
m = int(input())

# 5
# 1 2 4 2 5
# 3
# 1 - 2 * 4 /

pairs = []
splits = input().split()
for i in range(m):
    pairs.append((int(splits[2*i]), splits[2*i+1]))

origin = sum(arr)
for idx, new_symb in pairs:
    a, b = arr[idx-1], arr[idx]
    new_val = origin
    new_val -= a + b
    if new_symb == "-":
        new_val += (a - b)
    elif new_symb == "+":
        new_val += (a + b)
    elif new_symb == "*":
        new_val += a*b
    else:
        new_val += a/b
    print(new_val)


