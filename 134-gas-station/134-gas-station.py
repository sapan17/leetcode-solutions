class Solution(object):
    def canCompleteCircuit(self, gas, cost):
        """
        :type gas: List[int]
        :type cost: List[int]
        :rtype: int
        """
        
        if sum(gas) < sum(cost):
            return -1
        res = 0
        total = 0
        for i in range(len(cost)):
            total += (gas[i] - cost[i])
            if total < 0:
                res = i+1
                total = 0
        return res