

class Solution:
    def reviewAttendance(self , days: int) -> int:
        MOD = 1000000007
        def dfs(pre, prepre, pos):
            if pos == days:
                return 1
            c_list = ['P']
            if pre != 'O':
                c_list.append('O')
            if not (pre == 'D' and prepre == 'D'):
                c_list.append('D')
            cnt = 0
            for cur in c_list:
                cnt = (cnt + dfs(cur, pre, pos+1)) % MOD
            return cnt
        if days == 0:
            return 0
        elif days == 1:
            return 3
        else:
            cnt = 0
            cc_list = ['PO', 'PP', 'PD', 'OP', 'OD', 'DP', 'DO', 'DD']
            for cc in cc_list:
                cnt = (cnt + dfs(cc[0], cc[1], 2)) % MOD
            return cnt

print(Solution().reviewAttendance(3))