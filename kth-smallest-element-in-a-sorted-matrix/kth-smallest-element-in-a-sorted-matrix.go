func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    res := []int{}
    for row:=0; row <n; row++{
        for col:= 0; col <n; col++{
            res = append(res,matrix[row][col])
        }
    }
    sort.Ints(res)
    return res[k-1]
}