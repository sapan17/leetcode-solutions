class Solution(object):
    def maximumWealth(self, accounts):
        """
        :type accounts: List[List[int]]
        :rtype: int
        """
        maxwealth = 0
        for i in range(len(accounts)):
            totalwealth = sum(accounts[i])
            maxwealth = max(maxwealth, totalwealth)
        return maxwealth