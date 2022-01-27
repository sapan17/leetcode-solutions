class Solution(object):
    def trap(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        water, maxLeft, maxRight = 0, 0, 0
        l , r = 0, len(height)-1
        
        while l < r:
            if height[l] < height[r]:
                maxLeft = max(maxLeft, height[l])
                water += maxLeft - height[l]
                l += 1
            
            else:
                maxRight = max(maxRight, height[r])
                water += maxRight - height[r]
                r -= 1
        return water
        