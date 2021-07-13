func removeElement(nums []int, val int) int {
    size := len(nums)
    first, last := 0, size-1
    
    for first <= last{
        if nums[first] == val{
            nums[first],nums[last] = nums[last],nums[first]
            last--
        } else{
            first++
        }
    }
    return first
    
}