func strStr(haystack string, needle string) int {
    if len(needle) == 0{
        return 0
    }
    
    length := len(needle)
    for i:= 0; i< len(haystack)-length+1; i++{
        if haystack[i:i+length] == needle{
            return i
        }
    }
    return -1
}