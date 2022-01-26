class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        maxprofit = 0
        minnum = prices[0]
        
        for i in prices:
            if minnum > i:
                minnum = i
            if i - minnum > maxprofit:
                maxprofit = i - minnum
        return maxprofit