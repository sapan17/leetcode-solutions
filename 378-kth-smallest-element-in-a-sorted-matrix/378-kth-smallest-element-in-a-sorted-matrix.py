class Solution(object):
    def kthSmallest(self, matrix, k):
        """
        :type matrix: List[List[int]]
        :type k: int
        :rtype: int
        """
        res = []
        for i in range(len(matrix)):
            for j in range(len(matrix[0])):
                res.append(matrix[i][j])
        
        res.sort()
        return res[k-1]