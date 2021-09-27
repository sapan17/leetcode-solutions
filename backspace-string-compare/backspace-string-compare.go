func backspaceCompare(s string, t string) bool {
    m := len(s)
    n := len(t)
    st1 := []byte{}
    st2 := []byte{}
    for i:=0; i<m; i++{
        if s[i] == '#' && len(st1) > 0 {
            st1 = st1[:len(st1)-1]
        } else if s[i] != '#'{
            st1 = append(st1, s[i])
        }
    }
    for i:=0; i<n; i++{
        if t[i] == '#' && len(st2) > 0{
            st2 = st2[:len(st2)-1]
        } else if t[i] != '#'{
            st2 = append(st2, t[i])
        }
    }
    
    if len(st1) != len(st2){
        return false
    }
    for i:=0; i<len(st1); i++{
        if st1[i] != st2[i]{
            return false
        }
    }
    return true
}