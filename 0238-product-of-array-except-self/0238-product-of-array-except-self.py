class Solution(object):
    def productExceptSelf(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        arr1 = [1]*len(nums)
        arr2 = [1]*len(nums)
        
        for i in range(1,len(nums)):
            arr1[i] = arr1[i-1]*nums[i-1]
        for i in range(len(nums)-2,-1,-1):
            arr2[i] = arr2[i+1]*nums[i+1]
        for i in range(len(nums)):
            arr2[i] = arr1[i]*arr2[i]
        
        return arr2
