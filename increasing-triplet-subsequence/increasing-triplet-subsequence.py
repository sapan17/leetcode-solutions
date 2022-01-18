class Solution(object):
    def increasingTriplet(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        j = k = float('inf') 
        for i in range(len(nums)):
            if nums[i] <= j:
                j = nums[i]
            elif nums[i] <= k:
                k = nums[i]
            else:
                return True
        return False