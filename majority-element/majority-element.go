func majorityElement(nums []int) int {
    candidate, rating := nums[0], 1
    for _, n := range nums[1:]{
        if n == candidate{
            rating++
        } else{
            rating--
        }
        if rating == 0{
            candidate = n
            rating = 1
        }
    }
    return candidate
}