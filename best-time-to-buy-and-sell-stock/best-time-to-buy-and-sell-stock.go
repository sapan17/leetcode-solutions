func maxProfit(prices []int) int {
    n := len(prices)
    minprice := prices[0]
    maxprofit := 0
    for i:=1; i<n; i++{
        if prices[i] < minprice{
            minprice = prices[i]
        }
        if maxprofit < prices[i] - minprice{
            maxprofit = prices[i] - minprice
        }
    }
    return maxprofit
}