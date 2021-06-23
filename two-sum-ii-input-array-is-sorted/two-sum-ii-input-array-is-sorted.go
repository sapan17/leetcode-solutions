func twoSum(numbers []int, target int) []int {
    c1 := 0
    c2 := len(numbers)-1
    
    for c1 < c2 {
        sum := numbers[c1] + numbers[c2]
        if sum > target{
            c2--
        } else if sum < target{
            c1++
        } else{
            return []int {c1+1,c2+1}
        }
    }
    return []int{}
}