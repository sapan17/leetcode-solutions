func sortedSquares(nums []int) []int {
    
    n := len(nums)
    for i :=0; i< n; i++{
        nums[i] *= nums[i]
    }
    sort.Ints(nums)
    return nums  
}
