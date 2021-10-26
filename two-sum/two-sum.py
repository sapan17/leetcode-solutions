class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        preMap = {}
        
        for i, a in enumerate(nums):
            diff = target - a
            if diff in preMap:
                return [preMap[diff], i]
            preMap[a] = i
        return 