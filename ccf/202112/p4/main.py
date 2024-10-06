from re import split


class Node :
    def __init__(self, val=0, pre=None, pid=0):
        self.val = val
        self.pre = pre
        self.pid = pid

    def is_free(self, cur_pid=0):
        return self.pid == cur_pid or self.pid == 0

    def can_recover(self, pid):
        return self.pre and self.pre[1] == pid

    def delelte(self):
        self.pre = (self.val, self.pid)
        self.val = 0
        self.pid = 0

    def recover(self):
        self.val, self.pid = self.pre
        self.pre = None

    def set(self, x, pid):
        self.val = x
        self.pid = pid
        self.pre = None

    def get(self):
        return self.pid, self.val


class FS:
    def __init__(self, m):
        self.data = [Node() for _ in range(m)]

    def write(self, pid, l, r, x):
        for i in range(l, r+1):
            if not self.data[i].is_free(pid):
                return i-1 if i > l else -1
            self.data[i].set(x, pid)
        return r

    def delete_all(self, pid, l, r ):
        for i in range(l, r+1):
            if not self.data[i].pid == pid:
                return False

        for i in range(l, r+1):
            self.data[i].delelte()
        return True

    def recover_all(self, pid, l, r):
        for i in range(l, r+1):
            if not self.data[i].can_recover(pid):
                return False

        for i in range(l, r+1):
            self.data[i].recover()
        return True

    def read_one(self, p):
        return self.data[p].get()

n, m, k = map(int, input().split())
fs = FS(m+1)

for _ in range(k):
    splits = list(map(int, input().split()))
    opt = splits[0]
    if opt == 0:
        print(fs.write(splits[1], splits[2], splits[3], splits[4]))
    elif opt == 1:
        print("OK" if fs.delete_all(splits[1], splits[2], splits[3]) else "FAIL")
    elif opt == 2:
        print("OK" if fs.recover_all(splits[1], splits[2], splits[3]) else "FAIL")
    else:
        print(' '.join(map(str, fs.read_one(splits[1]))))