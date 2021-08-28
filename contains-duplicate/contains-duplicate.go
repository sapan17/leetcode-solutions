func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    n := len(nums)-1
    for i:=0; i<n; i++{
        if nums[i] == nums[i+1]{
            return true
        } 
    }
    return false
}