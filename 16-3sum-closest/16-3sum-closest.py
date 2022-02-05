class Solution(object):
    def threeSumClosest(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        nums.sort()
        diff = float('inf')
        for i in range(len(nums)-1):
            lo, hi = i+1, len(nums)-1
            while lo < hi:
                threesum = nums[i] + nums[lo] + nums[hi]
                
                if abs(target-threesum) < abs(diff):
                    diff = target - threesum

                if threesum > target:
                    hi -= 1

                if threesum < target:
                    lo += 1

                if threesum == target:
                    break
        return target - diff
        