class Solution(object):
    def increasingTriplet(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        nums1 = nums2 = float('inf')
        
        for i in nums:
            if i <= nums1:
                nums1 = i
            elif i <= nums2:
                nums2 = i
            else:
                return True
        return False