func maximumWealth(accounts [][]int) int {
    MaxWealth := 0
    for i :=0; i< len(accounts); i++{
        CurrentWealth := 0
        for j :=0; j < len(accounts[i]); j++{
            CurrentWealth += accounts[i][j]
        }
           if CurrentWealth > MaxWealth{
        MaxWealth = CurrentWealth
    } 
    }
    return MaxWealth
}