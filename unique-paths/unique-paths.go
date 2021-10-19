func uniquePaths(m int, n int) int {
    comb := make([][]int, m)
    for i := 0; i < m; i++ {
        comb[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 || j == 0 {
                comb[i][j] = 1
            } else {
                comb[i][j] = comb[i - 1][j] + comb[i][j - 1]
            }
        }
    }
    return comb[m - 1][n - 1]
}