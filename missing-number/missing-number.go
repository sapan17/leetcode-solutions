func missingNumber(nums []int) int {
    n := len(nums)
    c := (n*(n+1))/2
    e := 0
    for _, v:= range nums{
        e += v
    }
    return c - e
}
