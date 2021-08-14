func isPalindrome(s string) bool {
    l, r := 0, len(s)-1
    for r > l{
        if !isValid(s[l]){
            l++
            continue
        }
        if !isValid(s[r]){
            r--
            continue
        }
        if lower(s[l]) != lower(s[r]){
            return false
        }
        l++
        r--
    }
    return true
}

func lower(c byte)byte{
    if c >= 'A' && c <= 'Z'{
        return c + ' '
    }
    return c
}

func isValid(c byte)bool{
    if (c>= '0' && c<= '9') || (c>='A' && c<='Z') || (c>= 'a' && c<= 'z'){
        return true
    }
    return false
}