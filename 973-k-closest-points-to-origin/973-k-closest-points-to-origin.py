class Solution(object):
    def kClosest(self, points, k):
        """
        :type points: List[List[int]]
        :type k: int
        :rtype: List[List[int]]
        """
        points.sort(key=self.square)
        return points[:k]
    
    def square(self, point):
        return point[0] ** 2 + point[1] ** 2