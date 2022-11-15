class Solution(object):
    def containsDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        dictionary = {}
        for i in range(len(nums)):
            if nums[i] in dictionary:
                return True
            else:
                dictionary[nums[i]] = 1
        return False