func longestPalindrome(s string) string {
    if len(s) < 1{
        return ""
    }
    
    start := 0
    end := 0
    
    for i:=0; i<len(s); i++{
        len1 := expand(s, i, i)
        len2 := expand(s, i, i+1)
        len3 := max(len1,len2)
        if(len3 > end - start){
            start = i - ((len3-1)/2)
            end = i + (len3/2)
        }
    }
    return s[start:end+1]
}

func expand(s string, left, right int) int{
    if len(s) == 0 || left > right{
        return 0
    }
    for left >= 0 && right < len(s) && s[left] == s[right]{
        left -= 1
        right += 1
    }
    return right - left - 1
}

func max(a,b int) int{
    if a>b{
        return a
    } else{
        return b
    }
}