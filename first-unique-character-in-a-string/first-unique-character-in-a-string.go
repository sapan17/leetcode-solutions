func firstUniqChar(s string) int {
    n := len(s)
    arr := [26]int{}
    for i:=0; i<n; i++{
        arr[s[i]-97] += 1
    }
    for i := 0; i<n; i++{
        if arr[s[i]-97] == 1{
            return i
        }
    }
    return -1
}