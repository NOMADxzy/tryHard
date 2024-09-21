#
# 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
#
#
# @param m int整型 水池容量
# @param t int整型 总时长
# @param m1 int整型 进水阀打开时每分钟进水量
# @param t1 int整型 进水阀开关切换时长
# @param m2 int整型 排水阀打开时每分钟排水量
# @param t2 int整型 排水阀开关切换时长
# @return int整型 t时刻水池中的水量
#
class Solution:
    def dg(self, a, b):
        if a == b:
            return a
        if a<b:
            a,b = b,a
        mul_val = a * b
        while b!=0:
            a, b = b, a%b
        return mul_val // a

    def ComputeRemanentWater(self, m: int, t: int, m1: int, t1: int, m2: int, t2: int) -> int:
        cur_time = 0

        points = [(0, 0)]
        cur_t1, cur_t2 = 0, 0

        def is_open1(t):
            ts = t // t1
            if t%t1 == 0:
                ts += 1
            return ts % 2 == 0

        def is_open2(t):
            ts = t // t2
            if t%t2 == 0:
                ts += 1
            return ts % 2 == 0

        epoch_len = self.dg(2*t1, 2*t2)

        while cur_time < t:
            new_time = t
            if cur_t1 + t1 == cur_t2 + t2:
                new_time = min(new_time, cur_t1 + t1)
                delta = new_time - points[-1][0]
                cur_t1 += delta
                cur_t2 += delta
            elif cur_t1 + t1 < cur_t2 + t2:
                new_time = min(new_time, cur_t1 + t1)
                delta = new_time - points[-1][0]
                cur_t1 += delta
            else:
                new_time = min(new_time, cur_t2 + t2)
                delta = new_time - points[-1][0]
                cur_t2 += delta
            new_f = 0
            if is_open1(new_time):
                new_f += m1
            if is_open2(new_time):
                new_f -= m2
            points.append((new_time, new_f))
            cur_time = new_time

        cur_vol = 0
        for i in range(1, len(points)):
            delta_time = points[i][0] - points[i-1][0]
            f = points[i][1]
            new_vol = cur_vol + f * delta_time
            if new_vol<0:
                new_vol = 0
            if new_vol>m:
                new_vol = m
            cur_vol = new_vol

        return cur_vol

# s = Solution()
# print(s.ComputeRemanentWater(10,2,1,5,2,5))