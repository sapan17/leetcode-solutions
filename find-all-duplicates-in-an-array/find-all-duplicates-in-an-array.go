func findDuplicates(nums []int) []int {
    n := len(nums)-1
    var result []int
    sort.Ints(nums)
    for i:=0; i<n; i++{
        if nums[i] == nums[i+1]{
            result = append(result,nums[i])
        }
    }
    return result
}