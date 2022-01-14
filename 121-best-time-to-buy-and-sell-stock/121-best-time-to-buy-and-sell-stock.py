class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        minprice = prices[0]
        maxprofit = 0
        
        for i in prices:
            if i < minprice:
                minprice = i
            if i - minprice > maxprofit:
                maxprofit = i - minprice
        return maxprofit