func generate(numRows int) [][]int {
    res := [][]int{}
    res = append(res, []int{1})
    for i:= 2; i <= numRows; i++{
        temp := []int{}
        for j:= 0; j < i; j++{
            if j == 0 || j == i-1{
                temp = append(temp, 1)
            } else{
                temp = append(temp, res[i-2][j-1] + res[i-2][j])
            }
        }
        res = append(res, temp)
    }
    return res
}