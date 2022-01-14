class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        n = len(height)
        l = 0
        r = n-1
        MaxArea = 0
        while r > l:
            area = min(height[l], height[r]) * (r - l)
            MaxArea = max(MaxArea, area)
            
            if height[l] > height[r]:
                r -= 1
            else:
                l += 1
                
        return MaxArea
            