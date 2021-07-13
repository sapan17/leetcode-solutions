func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    
    letter := map[byte]int{}
    for i:= 0; i<len(s); i++{
        letter[s[i]]++
        letter[t[i]]--
    }
    
    for _, v := range letter{
        if v != 0{
            return false
        }
    }
    return true
}