class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        preMap = {}
        for i, a in enumerate(nums):
            if a in preMap:
                return True
            preMap[a] = i
        return False