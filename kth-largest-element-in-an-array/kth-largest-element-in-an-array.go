func findKthLargest(nums []int, k int) int {
    sort.Ints(nums)
    n := len(nums)
    return nums[n-k]
}