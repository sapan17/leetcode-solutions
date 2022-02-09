class Solution(object):
    def floodFill(self, image, sr, sc, newColor):
        """
        :type image: List[List[int]]
        :type sr: int
        :type sc: int
        :type newColor: int
        :rtype: List[List[int]]
        """
        m, n = len(image), len(image[0])
        oldcolor = image[sr][sc]
        
        if oldcolor == newColor:
            return image
        
        def dfs(i, j):
            if m > i >= 0 and n > j >= 0 and image[i][j] == oldcolor:
                image[i][j] = newColor
                
                dfs(i+1, j)
                dfs(i-1, j)
                dfs(i, j+1)
                dfs(i, j-1)
        
        dfs(sr, sc)    
        return image
                