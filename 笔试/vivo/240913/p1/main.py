#
# 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
#
# 获取最小分组数
# @param staff int整型一维数组 员工数组，其中0表示新员工，1表示老员工
# @return int整型

# 0 0 0
# 1 0
# 0 0
# 0
# 1
#
class Solution:
    def staffGroup(self , staff: List[int]) -> int:
        cnt0, cnt1 = 0, 0
        for e in staff:
            if e == 1:
                cnt1 += 1
            else:
                cnt0 += 1
        ans = 0
        if cnt1 > cnt0:
            ans = cnt1
        else:
            ans = cnt1
            cnt0 -= cnt1
            ans += cnt0 // 3
            if cnt0 % 3 > 0:
                ans += 1
        return ans