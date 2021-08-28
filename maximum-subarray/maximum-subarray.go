func maxSubArray(nums []int) int {
    if len(nums) == 1{
        return nums[0]
    }
    maxi := nums[0]
    curr := nums[0]
    for i:=1; i<len(nums);i++{
        curr = max(nums[i]+curr, nums[i])
        maxi = max(curr, maxi)
    }
    return maxi
}

func max(a,b int)int{
    if a > b{
        return a
    } else{
        return b
    }
}