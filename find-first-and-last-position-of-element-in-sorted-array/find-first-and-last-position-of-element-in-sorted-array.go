func searchRange(nums []int, target int) []int {
    return []int{findFirst(nums,target),findLast(nums,target)}
}
    
    func findFirst(nums []int, target int) int {
        left := 0
        right := len(nums)
        for left < right{
            mid := left + (right-left)/2
            if nums[mid] >= target{
                right = mid
            } else{
                left = mid +1
            }
        }
    
    if left < len(nums) && nums[left] == target{
        return left
    } else {
        return -1
    }
}
    func findLast(nums []int, target int) int{
        left := 0
        right := len(nums)-1
        for left < right{
            mid := (left + right + 1)/2
            if nums[mid] <= target{
                left = mid
            } else{
                right = mid -1
            }
        }
        if left < len(nums) && nums[left] == target{
            return left
        } else {
            return -1
        }
}
   
