func longestConsecutive(nums []int) int {
    if len(nums) < 2{
        return len(nums)
    }
    
    longest := 1
    current := 1
    
    sort.Ints(nums)
    
    for i:=1; i<len(nums); i++{
        if nums[i] != nums[i-1]{
            if nums[i] == nums[i-1] + 1{
                current +=1
            } else{
                longest = max(longest,current)
                current = 1
            }
        } 
    }
    return max(longest, current)
}

func max(a,b int) int{
    if a > b{
        return a
    } else{
        return b
    }
}