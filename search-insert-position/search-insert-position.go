func searchInsert(nums []int, target int) int {
    l := 0
    r := len(nums)
    for l < r{
        mid := (l+r)/2
        if target == nums[mid]{
            return mid
        }
        if target > nums[mid]{
            l = mid + 1
        }else{
            r = mid - 1
        }
    }
    if l < len(nums) && target > nums[l]{
        return l+1
    }
    return l
}