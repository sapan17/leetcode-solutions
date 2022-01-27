class Solution(object):
    def countPrimes(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n <= 2:
            return 0
        
        
        number = [False, False] + [True] * (n-2)
        for p in range(2, int(sqrt(n))+1):
            if number[p]:
                for multiple in range(p*p,n,p):
                    number[multiple] = False
        
        return sum(number)