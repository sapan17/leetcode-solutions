class Solution(object):
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        
        maxSub = nums[0]
        cursum = 0
        
        for i in nums:
            if cursum < 0:
                cursum = 0
            cursum += i
            maxSub = max(maxSub, cursum)
        return maxSub