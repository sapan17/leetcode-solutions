func calculate(s string) int {
    ans, last, cur := 0, 0, 0
    op := '+'
    for i,ch := range s{
        if ch-'0' >= 0 && ch-'0' <= 9{
            cur =  cur*10 + int(ch-'0')
        }
       if ch == '+' || ch == '-' || ch =='*' || ch =='/' || i == len(s)-1{
            switch op{
            case '+':
                ans += last
                last = cur
            case '-':
                ans += last
                last = -cur
            case '*':
                last *= cur
            case '/':
                last /= cur
        }
        op = ch
        cur = 0
    }
}
    ans += last
    return ans
}