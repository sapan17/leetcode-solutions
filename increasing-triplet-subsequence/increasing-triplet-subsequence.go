func increasingTriplet(nums []int) bool {
    l := 10000000000
    r := 10000000000
    
    for _,num := range nums{
        if num <= l{
            l = num
        } else if num <= r{
            r = num
        } else{
            return true
        }
    }
    return false
}