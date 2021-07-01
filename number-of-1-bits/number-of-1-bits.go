func hammingWeight(num uint32) int {
    res := 0
    for num != 0{
        res += int(num%2)
        num /= 2
    }
    return res
}