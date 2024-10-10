from typing import List
# moneyArrays包含5，10，20的付款序列，商品价格为5，能否全部购买成功
class Solution:
    def soldSolution(self , moneyArrays: List[int]) -> bool:
        # write code here
        cnt5, cnt10, cnt20 = 0, 0, 0
        for money in moneyArrays:
            if money == 5:
                cnt5 += 1
            elif money == 10:
                if cnt5 == 0:
                    return False
                cnt10 += 1
                cnt5 -= 1
            else:
                if cnt10>0:
                    if cnt5==0:
                        return False
                    cnt10 -= 1
                    cnt5 -= 1
                else:
                    if cnt5<3:
                        return False
                    cnt5 -= 3
        return True

