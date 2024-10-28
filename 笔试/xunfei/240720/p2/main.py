
n = int(input())
prices = list(map(int, input().split()))

prices.sort()
# 5
# 1 2 3 4 5

ans = []
l, r = n//2, n//2
if n%2==0:
    l -= 1
else:
    r += 1
while l>=0 and r<n:
    # if l<0:
    #     ans.extend(prices[r:])
    #     break
    # if r>=n:
    #     tmp = prices[:l+1]
    #     tmp.reverse()
    #     ans.extend(tmp)
    #     break
    if (l+1+n-r) % 2 == 1:
        if l+1 > n-r:
            ans.append(prices[l])
            l -= 1
        else:
            ans.append(prices[r])
            r += 1
    else:
        if prices[l] <= prices[r]:
            ans.append(prices[l])
            l -= 1
        else:
            ans.append(prices[r])
            r += 1

if l==0:
    ans.append(prices[l])
if r==n-1:
    ans.append(prices[r])

print(*ans)
