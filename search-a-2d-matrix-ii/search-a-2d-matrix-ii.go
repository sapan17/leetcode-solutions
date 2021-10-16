func searchMatrix(matrix [][]int, target int) bool {
    n := len(matrix)
    m := len(matrix[0])
    for row:=0; row< n; row++{
        for col :=0; col <m; col++{
            if matrix[row][col] == target{
                return true
            }
        }
    }
    return false
}