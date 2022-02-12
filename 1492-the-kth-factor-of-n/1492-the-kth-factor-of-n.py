class Solution(object):
    def kthFactor(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: int
        """
        divisor, sqrt_n = [], int(n**0.5)
        for i in range(1, sqrt_n+1):
            if n % i == 0:
                k -= 1
                divisor.append(i)
                
                if k == 0:
                    return i
        if (sqrt_n * sqrt_n) == n:
            k += 1
            
        div_len = len(divisor)
        return n // divisor[div_len - k] if k <= div_len else -1