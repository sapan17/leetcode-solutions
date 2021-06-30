func reverseBits(num uint32) uint32 {
    var res uint32
    for i := 0; i< 32; i++{
        res *= 2
        if num != 0{
            res += num %2
            num /= 2
        }
    }
    return res
}