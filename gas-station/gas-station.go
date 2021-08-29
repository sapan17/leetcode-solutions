func canCompleteCircuit(gas []int, cost []int) int {
    diff_total := 0
    start, tank := 0,0
    for i:=0; i<len(gas);i++{
        diff_total += gas[i] - cost[i]
    }
    if diff_total < 0{
        return -1
    } else{
        for i:=0; i<len(gas);i++{
            tank = tank + gas[i] - cost[i]
            if tank < 0{
                start = i+1
                tank = 0
            }
        }
    }
    return start
}
