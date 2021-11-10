class Solution(object):
    def productExceptSelf(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        n = len(nums)
        nums1 = [1] * n
        nums2 = [1] * n
        
        for i in range(1, n):
            nums1[i] = nums[i-1] * nums1[i-1]
        for i in range(n-2, -1, -1):
            nums2[i] = nums[i+1] * nums2[i+1]
        for i in range(n):
            nums2[i] = nums1[i] * nums2[i]
            
        return nums2