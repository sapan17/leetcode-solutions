func isHappy(n int) bool {
    slow := findSquareSum(n)
    fast := findSquareSum(findSquareSum(n))
    
    for slow != fast{
        slow = findSquareSum(slow)
        fast = findSquareSum(findSquareSum(fast))
    }
    return slow == 1
}

func findSquareSum(a int)int {
    sum := 0
    digit := 0
    for a > 0{
        digit = a%10
        sum += digit * digit
        a /= 10
    }
    return sum
}