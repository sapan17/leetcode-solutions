func romanToInt(s string) int {
    z := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    var res int
    for i := 0; i < len(s); i++{
        if i != len(s)-1 && z[s[i]] < z[s[i+1]]{
            res -= z[s[i]]
        } else{
            res += z[s[i]]
        }
    }
    return res
}