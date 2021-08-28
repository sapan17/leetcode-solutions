func climbStairs(n int) int {
    if n <= 3 {
        return n
    }
    prev, curr := 1, 2
    for i:= 3; i<=n; i++ {
        prev, curr = curr, prev+curr
    }
    return curr
}