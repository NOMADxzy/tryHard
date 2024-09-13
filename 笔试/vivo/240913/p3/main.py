from typing import List
#
# 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
#
#
# @param prices int整型一维数组 礼物的价格
# @param k int整型 粉丝人数
# @return bool布尔型
# 3 3 4 2 2 1
class Solution:
    def canEqualDistribution(self, prices: List[int], k: int) -> bool:
        total = sum(prices)
        if total % k != 0:
            return False
        target = total // k
        price_cnt = {}
        max_price = 0
        for p in prices:
            if p > target:
                return False
            if p not in price_cnt:
                price_cnt[p] = 0
            price_cnt[p] += 1
            max_price = max(max_price, p)

        def dfs(f: int, cnt: int, t: int) -> bool:
            if f==t:
                return True
            for p in range(t-f, 0, -1):
                if cnt == 0:
                    break
                if p in price_cnt and price_cnt[p] > 0:
                    cur_cnt = min(cnt, price_cnt[p])
                    price_cnt[p] -= cur_cnt
                    if dfs(p, cur_cnt, t-f): # 贪心的从大到小递归尝试
                        cnt -= cur_cnt
                        return True
                    else:
                        price_cnt[p] += cur_cnt

            return False
        for p in range(max_price, 0, -1):
            if p in price_cnt and price_cnt[p] > 0:
                while price_cnt[p] > 0:
                    price_cnt[p] -= 1
                    if not dfs(p, 1, target):
                        return False

        return True





print(Solution().canEqualDistribution([4, 3, 3, 2, 2, 2], 2))