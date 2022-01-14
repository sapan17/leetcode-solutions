class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        count = {}
        
        for i, a in enumerate(nums):
            res = target - a
            
            if res in count:
                return [i, count[res]]
            count[a] = i
        return