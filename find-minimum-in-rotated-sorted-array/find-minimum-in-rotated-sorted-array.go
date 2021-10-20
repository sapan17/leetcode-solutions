func findMin(nums []int) int {
    if len(nums) < 2{
        return nums[0]
    }
    l := 0
    r := len(nums)-1
    if nums[r] > nums[0]{
        return nums[0]
    }
    for r >= l{
        mid := l +(r-l)/2
        if nums[mid] > nums[mid+1]{
            return nums[mid+1]
        }
        if nums[mid-1] > nums[mid]{
            return nums[mid]
        }
        if nums[mid] > nums[0]{
            l = mid+1
        } else{
            r = mid-1
        }
    }
    return -1
}