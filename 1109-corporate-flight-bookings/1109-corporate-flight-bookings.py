class Solution(object):
    def corpFlightBookings(self, bookings, n):
        """
        :type bookings: List[List[int]]
        :type n: int
        :rtype: List[int]
        """
        res = [0] * (n+1)
        result = []
        dummy = 0
        for start, end, cost in bookings:
            res[start-1] += cost
            res[end] -= cost
        
        for i in range(len(res)-1):
            dummy += res[i]
            result.append(dummy)
        return result