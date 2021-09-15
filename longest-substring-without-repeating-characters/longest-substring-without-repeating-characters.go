func lengthOfLongestSubstring(s string) int {
    left, right := 0,0
    longest := 0
    m := make([]bool, 256)
    
    for right < len(s){
        for left < right &&  m[s[right]]{
            m[s[left]] = false
            left += 1
        }
        
        longest = max(right-left +1, longest)
        m[s[right]] = true
        right += 1
    }
    return longest
}

func max(a, b int) int{
    if a >= b{
        return a
    } else{
        return b
    }
}