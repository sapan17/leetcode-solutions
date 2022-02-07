class Solution(object):
    def prisonAfterNDays(self, cells, n):
        """
        :type cells: List[int]
        :type n: int
        :rtype: List[int]
        """
        n = (n-1) % 14 + 1
        
        while n:
            temp = [0] * 8
            for i in range(1,7):
                if cells[i-1] == cells[i+1]:
                    temp[i] = 1
                else:
                    temp[i] = 0
            cells = temp
            n -= 1
        return cells
            