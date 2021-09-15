func productExceptSelf(nums []int) []int {
    n := len(nums)
    res1 := make([]int, n)
    res2 := make([]int, n)
    res3 := make([]int, n)
    res1[0] = 1
    res2[n-1] = 1
    for i := 1; i < n; i++ {
        res1[i] = res1[i-1] * nums[i-1]
    }
    for i:= n-2; i >= 0; i--{
        res2[i] = res2[i+1] * nums[i+1]
    }
    for i := 0; i < n; i++ {
        res3[i] = res1[i]*res2[i]
    }
    return res3
}