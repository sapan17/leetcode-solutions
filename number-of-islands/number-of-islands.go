var m,n int

func numIslands(board [][]byte) int {
    m = len(board)
    n = len(board[0])
    if m == 0 || n == 0 {
        return 0
    }
    res := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == '1' {
                res += 1
                dfs(&board, i, j) // pass the reference of the whole board
            }
        }
    }
    
    return res
    
}

func dfs(board *[][]byte, i, j int) {
    if i < 0 || i >= m || j < 0 || j >= n || (*board)[i][j] == '0' {
        return
    }
    
    (*board)[i][j] = '0' // dereference of the pointer and you will get the original board
    dfs(board, i+1, j)
    dfs(board, i, j-1)
    dfs(board, i, j+1)
    dfs(board, i-1, j)
}