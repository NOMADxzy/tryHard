n = int(input())
arr = list(map(int, input().split()))
# 6
# 1 1 2 3 4 5

stack = []
for v in arr:
    stack.append(v)
    while len(stack)>1 and stack[len(stack)-1] == stack[len(stack)-2]:
        new_val = stack[len(stack)-1] + 1
        stack = stack[:-2]
        stack.append(new_val)

print(" ".join(map(str, stack)))