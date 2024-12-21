
n, size = map(int, input().split())

def transfer_to_int(s: str):
    ret = 0
    mask = 1
    for i in range(len(s)-1, -1, -1):
        if s[i] != '0':
            ret += (ord(s[i]) - 48) * mask
        mask *= 10
    return ret

t_queue = list(map(transfer_to_int, input().split()))
m_queue = list(map(transfer_to_int, input().split()))

exists = set(t_queue)
tmp_set = set(t_queue[:size])
ptr = size
ans = 0

while m_queue:
    ans += 1
    m = m_queue[0]
    m_queue = m_queue[1:]
    if m not in exists:
        continue
    if m in tmp_set:
        tmp_set.remove(m)
        if ptr<len(t_queue):
            tmp_set.add(t_queue[ptr])
            ptr += 1
    else:
        m_queue.append(m)

print(ans)

