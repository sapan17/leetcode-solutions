class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        premap = {}
        result = 0
        res = []
        for i, a in enumerate(nums):
            result = target - a
            if result in premap:
                return[premap[result], i]
            premap[a] = i
        return