# 给一个字符串，第i(0=<i<n)次将s[i]移到末尾

s = input()
n = len(s)

class ListNode:
    def __init__(self, val, next):
        self.val, self.next = val, next

head = ListNode(s[0],None)
tail = head

for i in range(1,n):
    tail.next = ListNode(s[i], None)
    tail = tail.next

cur = head
ans = []
while True:

    ans.append(cur.next.val)
    tmp = cur.next.next
    if tmp == None:
        ans.append(cur.val)
        break
    cur.next = None
    tail.next = cur
    tail = tail.next
    cur = tmp

print("".join(ans))
