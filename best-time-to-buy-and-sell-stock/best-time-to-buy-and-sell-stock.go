func maxProfit(prices []int) int {
    if len(prices) < 2{
        return 0
    }
    minprice := prices[0]
    maxprofit := 0
    for i:=0;i<len(prices);i++{
        if prices[i]< minprice{
            minprice = prices[i]
        } else if prices[i] - minprice > maxprofit{
            maxprofit = prices[i] - minprice
        }
    }
    return maxprofit
}