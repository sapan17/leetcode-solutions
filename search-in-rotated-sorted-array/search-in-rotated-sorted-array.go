func search(nums []int, target int) int {
    l := 0
    r := len(nums)-1
    for l <= r{
        mid := (l+r)/2
        if target == nums[mid]{
            return mid
        } else if nums[mid] >= nums[l]{
            if nums[mid] >= target && target >= nums[l]{
                r = mid - 1
            } else{
                l = mid + 1
            }
        } else{
            if nums[mid] <= target && nums[r]>= target{
                l = mid + 1
            } else{
                r = mid - 1
            }
        }
    }
    return -1
}