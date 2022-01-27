class Solution(object):
    def maxProduct(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        maxval, minval = 1, 1
        res = max(nums)
        for i in nums:
            if i == 0:
                maxval, minval = 1, 1
            tmp = maxval * i
            maxval = max(maxval*i, minval*i, i)
            minval = min(tmp, minval*i, i)
            res = max(res, maxval)
        return res