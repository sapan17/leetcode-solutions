func myPow(x float64, n int) float64 {
    if n == 0{
        return 1
    }
    if x == 0{
        return 0
    }
    if n < 0{
        n = -n
        x = 1/x
    }
    res := 1.0
    for n > 0{
        if n%2 == 1{
            res *= x
        }
        x *= x
        n >>= 1
    }
    return res
}