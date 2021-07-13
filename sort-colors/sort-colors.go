func sortColors(nums []int)  {
    if len(nums) <= 1{
        return
    }
    l := 0
    r := len(nums)-1
    for i:=0; i<=r; i++{
        if nums[i] == 2{
            nums[r],nums[i] = nums[i],nums[r]
            r--
            i--
        }else if nums[i] == 0{
            nums[l],nums[i] = nums[i],nums[l]
            l++
        }
        
    }
}