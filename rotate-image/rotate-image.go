func rotate(matrix [][]int)  {
    transpose(matrix)
    reflect(matrix)
}

func transpose(matrix [][]int){
    n := len(matrix)
    for i:=0; i<n; i++{
        for j:=i; j<n; j++{
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
}

func reflect(matrix [][]int){
    n := len(matrix)
    for i:=0; i<n; i++{
        for j:=0; j<n/2; j++{
            matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
        }
    }
}