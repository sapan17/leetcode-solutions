func spiralOrder(matrix [][]int) []int {
    var res []int
    n := len(matrix)
    m := len(matrix[0])
    left,right,up,down := 0, m-1, 0, n-1
    for len(res) < m*n{
        for i:=left; i<=right && len(res)<n*m; i++{
            res = append(res,matrix[up][i])
        }
        for j:=up+1; j<=down-1 && len(res)<n*m; j++{
            res = append(res,matrix[j][right])
        }
        for i:=right; i>=left && len(res)<n*m; i--{
            res = append(res,matrix[down][i])
        }
        for j:=down-1; j>=up+1 && len(res)<n*m; j--{
            res = append(res,matrix[j][left])
        }
        left++
        right--
        up++
        down--
    }
    return res
}