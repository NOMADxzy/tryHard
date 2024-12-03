# 3 2
# The tHe thE
# the THE

m, n = map(int, input().split())
word_list1 = input().split()
word_list2 = input().split()
for i in range(m):
    word_list1[i] = word_list1[i].lower()
for i in range(n):
    word_list2[i] = word_list2[i].lower()

set1, both, only2 = set(), set(), set()

for w in word_list1:
    set1.add(w)
for w in word_list2:
    if w in set1:
        both.add(w)
    else:
        only2.add(w)

print(len(both), len(set1)+len(only2))
