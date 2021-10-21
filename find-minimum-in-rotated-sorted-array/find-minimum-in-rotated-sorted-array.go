func findMin(nums []int) int {
    l := 0
    r := len(nums) - 1
    
    if len(nums) < 2{
        return nums[0]
    }
    if nums[0] < nums[r]{
        return nums[0]
    }
    for l <= r{
        mid := (l+r)/2
        if nums[mid] > nums[mid+1]{
            return nums[mid+1]
        }
        if nums[mid] < nums[mid-1]{
            return nums[mid]
        }
         if nums[mid] > nums[l]{
            l = mid
         } else{
             r = mid
         }
    }
    return 0
}