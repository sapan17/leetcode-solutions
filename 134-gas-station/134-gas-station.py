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
        result = 0
        for i in range(len(gas)):
            res += gas[i] - cost[i]
            if res < 0:
                result = i+1
                res = 0
        return result
                