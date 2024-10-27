# ? 表示匹配0个或多个字符，问主串中能被模式串匹配上多少个子字符串

# 12123454567878
# 12?34?78


main_str = input()
sub_str = input()

m,n = len(main_str), len(sub_str)

tmp = []
for i in range(n):
    if sub_str[i] == '?':
       tmp.append(i)
assert len(tmp) == 2
i1, i2 = tmp

l1,l2,l3 = i1, i2-i1-1, n-i2-1

list1, list2, list3 = [], [], []
for i in range(m):
    if i+l1<=m and main_str[i:i+l1] == sub_str[:i1]:
        list1.append(i)
    if i+l2<=m and main_str[i:i+l2] == sub_str[i1+1:i2]:
        list2.append(i)
    if i+l3<=m and main_str[i:i+l3] == sub_str[i2+1:]:
        list3.append(i)

detail = []
ans = 0

for i in list1:
    for j in list2:
        if i+l1 > j:
            continue
        for idx, k in enumerate(list3):
            if j+l2 > k:
                continue
            else:
                detail.append((i,j,len(list3)-idx))
                ans += len(list3)-idx
                break

print(ans)