
s = input()
# 消消乐，给一串字符实现2个以上相邻的字符被消除
ans = ""
i = 0
while i<len(s):
    j = i
    while j<len(s) and s[j] == s[i]:
        j += 1
    j -= 1
    if len(ans)>0 and s[i] == ans[-1]:
        ans = ans[:-1]
    elif j == i:
        ans += s[i]
    i = j+1

print(ans if len(ans)>0 else "null")