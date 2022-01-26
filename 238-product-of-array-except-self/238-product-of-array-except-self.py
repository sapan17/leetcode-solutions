class Solution(object):
    def productExceptSelf(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        nums1 = [1]*len(nums)
        nums2 = [1]*len(nums)
        
        for i in range(1, len(nums)):
            nums1[i] = nums[i-1] * nums1[i-1]
        
        for i in range(len(nums)-2, -1, -1):
            nums2[i] = nums[i+1] * nums2[i+1]
            
        for i in range(len(nums)):
            nums2[i] = nums1[i] * nums2[i]
        
        return nums2