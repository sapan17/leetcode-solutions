class Solution(object):
    def subArrayRanges(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        res = 0
        n = len(nums)
        
        for i in range(n):
            mx = nums[i]
            mn = nums[i]
            
            for j in range(i+1, n):
                if nums[j] > mx:
                    mx = nums[j]
                elif nums[j] < mn:
                    mn = nums[j]
                res += (mx-mn)
        return res
            