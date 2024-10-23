# 8
# 3 7 2 1 6 5 4 8

# 1~n 的n个数的数组，最多可以

n = int(input())
arr = list(map(int, input().split()))

ans = []

pos = 0
cnts = [2] * (n+1)

while len(ans)<n:
    if cnts[arr[pos]] == 0:
        ans.append(arr[pos])
        pos += 1
    elif pos+1 == n:
        ans.append(arr[pos])
        break
    elif cnts[arr[pos+1]] == 0:
        ans.append(arr[pos])
        ans.append(arr[pos+1])
        pos += 2
        break
    elif pos+2==n:
        ans.append(arr[pos])
        ans.append(arr[pos + 1])
        if ans[-1] > ans[-2]:
            ans[-1], ans[-2] = ans[-2], ans[-1]
    else:
        if arr[pos+2] > arr[pos] and arr[pos+2] > arr[pos+1]:
            cnts[arr[pos+2]] -= 2
            cnts[arr[pos+1]] -= 1
            cnts[arr[pos]] -= 1
            arr[pos], arr[pos+1], arr[pos+2] = arr[pos+2], arr[pos], arr[pos+1]
            ans.append(arr[pos])
            pos += 1
        elif arr[pos+1] > arr[pos]:
            cnts[arr[pos + 1]] -= 1
            cnts[arr[pos]] -= 1
            arr[pos], arr[pos + 1] = arr[pos + 1], arr[pos]
            ans.append(arr[pos])
            pos += 1
        else:
            ans.append(arr[pos])
            pos += 1

print(" ".join(map(str, ans)))
