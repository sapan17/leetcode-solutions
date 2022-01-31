class Solution(object):
    def findMaxConsecutiveOnes(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        res = result = 0
        for i in range(len(nums)):
            if nums[i] == 1:
                res += 1
            else:
                res = 0
            result = max(res, result)
        return result