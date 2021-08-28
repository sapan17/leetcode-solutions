func max(a int,b int)int{
    if a > b{
        return a
    } else{
        return b
    }
}

func min(a int, b int)int{
    if a > b{
        return b
    } else{
        return a
    }
}

func maxProduct(nums []int) int {
    mini, maxi, res := nums[0], nums[0], nums[0]
    for i:= 1; i< len(nums); i++{
        if nums[i] < 0{
            temp := maxi
            maxi = mini
            mini = temp
        }
        mini = min(nums[i], nums[i]*mini)
        maxi = max(nums[i], nums[i]*maxi)
        res = max(maxi,res)
    }
    return res
}