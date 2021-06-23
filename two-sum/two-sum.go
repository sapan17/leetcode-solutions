func twoSum(nums []int, target int) []int {
    z := map[int]int{}
    for i, n:= range nums{
        if x, y := z[target-n]; y{
            return []int{x,i}
        }
        z[n] = i
    }
    return nil
} 