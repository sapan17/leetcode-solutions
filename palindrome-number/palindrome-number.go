func isPalindrome(x int) bool {
    if x < 0 || (x%10 == 0 && x != 0){
        return false
    }
    revert := 0
    for x > revert{
        revert = revert * 10 + x%10
        x /= 10
    }
    return x == revert || x == revert/10
}