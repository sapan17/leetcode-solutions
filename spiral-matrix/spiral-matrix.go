func spiralOrder(matrix [][]int) []int {
    var res []int
    r := len(matrix[0])
    c := len(matrix)
    
    left, right, up, down := 0, r -1, 0, c-1
    
    for len(res) < r*c{
        for i:=left; i<= right && len(res) < r*c; i++{
            res = append(res, matrix[up][i])
        }
        for j:=up+1; j<= down-1 && len(res) < r*c; j++{
            res = append(res, matrix[j][right])
        }
        for i:= right; i>= left && len(res) < r*c; i--{
            res = append(res, matrix[down][i])
        }
        for j := down-1; j>=up+1 && len(res) < r*c; j--{
            res = append(res, matrix[j][left])
        }
        left++
        right--
        up++
        down--
    }
    return res
}