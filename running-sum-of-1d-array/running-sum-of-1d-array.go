func runningSum(nums []int) []int {
    presum := 0
    for i := 0; i< len(nums); i++{
        presum = presum + nums[i]
        nums[i] = presum
    }
    return nums
}