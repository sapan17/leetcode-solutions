func max(c,d int)int{
    if c> d{
        return c
    } else{
        return d
    }
}

func rob(nums []int) int {
    a,b := 0,0
    for _,num := range nums{
        a,b = b,max(b,a+num) 
    }
    return b
}