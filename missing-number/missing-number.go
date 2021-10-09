func missingNumber(nums []int) int {
    n := len(nums)
    c := (n*(n+1))/2
    e := 0
    for i:=0; i<len(nums); i++{
        e += nums[i]
    }
    return c - e
}
