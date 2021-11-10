class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        minprices = prices[0]
        maxprofit = 0
        
        for i in range(len(prices)):
            if minprices > prices[i]:
                minprices = prices[i]
            if maxprofit < prices[i] - minprices:
                maxprofit = prices[i] - minprices
        return maxprofit