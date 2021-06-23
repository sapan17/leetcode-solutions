func numIdenticalPairs(nums []int) int {
    m := make(map[int]int)
    var pairs int
    for _, num := range nums {
        pairs += m[num]
        m[num]++
    }
    return pairs
}
