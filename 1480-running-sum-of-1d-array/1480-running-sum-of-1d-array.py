class Solution(object):
    def runningSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        res = []
        result = 0
        for i in nums:
            result += i
            res.append(result)
        return res