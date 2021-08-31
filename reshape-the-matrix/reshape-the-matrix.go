func matrixReshape(mat [][]int, r int, c int) [][]int {
    m := len(mat)
    n := len(mat[0])
    
    sum := m*n
    if sum != r * c{
        return mat
    }
    
    var res[][]int
    i := 0
    for a:=0; a< r && i<sum; a++{
        row := make([]int,c)
        for b := 0; b < c && i < sum; b++{
            row[b] = mat[i/n][i%n]
            i++
        }
        res = append(res,row)
    }
    return res
}