func maxProfit(prices []int) int {
    if len(prices) < 2{
        return 0
    }
    maxprofit := 0
    for i:=0;i<len(prices)-1;i++{
        if prices[i+1] > prices[i]{
            maxprofit += prices[i+1] - prices[i] 
        }
    }
    return maxprofit
}