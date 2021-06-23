func plusOne(digits []int) []int {
    idx := len(digits)-1
    for idx >=0{
        if digits[idx] != 9{
            digits[idx] += 1
            return digits
        } else{
            digits[idx] = 0
        }
        idx --
    }
    
    ans := []int{1}
    ans = append(ans, digits...)
    return ans
}