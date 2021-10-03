func solve(board [][]byte)  {
    
    if len(board) < 2 || len(board[0]) < 2{
        return
    }
    for i:=0; i< len(board); i++{
        dfs(board, i , 0)
        dfs(board, i, len(board[i])-1)
    }
    
    for j:=0; j<len(board[0]); j++{
        dfs(board, 0, j)
        dfs(board, len(board)-1, j)
    }
    
    for i:=0 ; i< len(board); i++{
        for j:=0; j<len(board[i]); j++{
            if board[i][j] == 'B'{
                board[i][j] = 'O'
            } else{
                board[i][j] = 'X'
            }
        }
    }
}

func dfs(board[][]byte, i, j int){
    if i<0 || i >= len(board) || j<0 || j >=len(board[i]) || board[i][j] != 'O'{
        return 
    }
    board[i][j] = 'B'
    
    if i - 1 >=0{
        dfs(board,i-1,j)
    }
    if j + 1 < len(board[i]){
        dfs(board,i,j+1)
    }
    if i + 1 < len(board){
        dfs(board, i+1, j)
    }
    if j - 1 >= 0{
        dfs(board, i, j-1)
    }
}