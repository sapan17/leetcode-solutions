func pivotIndex(nums []int) int {
    
    sum := 0
    n := len(nums)
    for i:=0; i<n; i++{
        sum += nums[i]
    }
    
    leftsum := 0
    for i:=0; i<n; i++{
        if leftsum == sum - nums[i] - leftsum{
            return i
        } else{
            leftsum += nums[i]
        }
    }
    return -1
}