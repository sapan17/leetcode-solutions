class Solution(object):
    def pivotIndex(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        sum_right = sum(nums)
        sum_left = 0
        
        for i in range(len(nums)):
            
            sum_right -= nums[i]
            
            if sum_right == sum_left:
                return i
            
            sum_left += nums[i]
        return -1