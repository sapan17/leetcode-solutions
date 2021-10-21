func missingNumber(nums []int) int {
    n := len(nums)
    c := (n*(n+1))/2
    for i:=0; i<len(nums); i++{
        c -= nums[i]
    }
    return c
}
