
n = int(input())
s = str(input())

# 01001011111010
ans = 1
l,r = 0,0
while r<n:
    if r>0 and s[r] == s[r-1]:
        l = r
    ans = max(ans, r-l+1)
    r += 1

if s[0] == s[n-1]:
    print(ans)
else:
    splited_ans = 0
    for i in range(0, n):
        if i>0 and s[i] == s[i-1]:
            splited_ans += i
            break
    for j in range(n-1, splited_ans-1, -1):
        if j<n-1 and s[i] == s[i+1]:
            break
        splited_ans += 1
    print(max(ans, splited_ans))



# n = int(input())
# s = str(input())
#
# # 01001011111010
#
# if s[0] == s[n-1]:
#     ans = 1
#     l,r = 0,0
#     while r<n:
#         if r>0 and s[r] == s[r-1]:
#             l = r
#         ans = max(ans, r-l+1)
#         r += 1
#     print(ans)