func removeDuplicates(nums []int) int {
    if len(nums)<=1{
        return len(nums)
    }
    var i,j int = 0,1
    for j < len(nums){
        if nums[i] != nums[j]{
            i++
            nums[i] = nums[j]
        }
        j++
    }
    return i +1
}
    
