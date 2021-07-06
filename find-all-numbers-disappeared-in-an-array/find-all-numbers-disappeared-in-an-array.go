func findDisappearedNumbers(nums []int) []int {
    lens := len(nums)
    j := 0
    if lens == 0{
        return nil
    }
    res := make([]int, lens)
    for _, v := range nums{
        res[v-1] = v
    }
    for i:=0; i< lens;i++{
        if res[i] == 0{
            res[j] = i + 1
            j++
        }
    }
    return res[:j]
}