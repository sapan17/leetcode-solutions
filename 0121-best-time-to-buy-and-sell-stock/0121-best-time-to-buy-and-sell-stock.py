class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        maxprofit = 0
        least = prices[0]
        
        for i in range(len(prices)):
            if prices[i] < least:
                least = prices[i]
            if prices[i] - least > maxprofit:
                maxprofit = prices[i]-least
        return maxprofit