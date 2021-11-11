class Solution(object):
    def lengthOfLIS(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        
        List = [1] * len(nums)
        
        for i in range(len(nums)-1, -1, -1):
            for j in range(i, len(nums)):
                if nums[i] < nums[j]:
                    List[i] = max(List[i], 1+List[j])
        return max(List)