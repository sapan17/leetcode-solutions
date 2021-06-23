func isPowerOfTwo(n int) bool {
    for n >= 2{
        if n%2 != 0{
            return false
        }
        n /= 2
    }
    return n==1
}