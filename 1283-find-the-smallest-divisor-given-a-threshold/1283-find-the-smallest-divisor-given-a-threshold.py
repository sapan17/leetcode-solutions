class Solution(object):
    def smallestDivisor(self, nums, threshold):
        """
        :type nums: List[int]
        :type threshold: int
        :rtype: int
        """
        l, r = 1, max(nums)
        
        while l < r:
            mid = (l + r) / 2
            if sum((num+mid-1)/mid for num in nums) > threshold:
                l = mid + 1
            else:
                r = mid
        return l
            