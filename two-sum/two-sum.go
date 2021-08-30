func twoSum(nums []int, target int) []int {
    hashTable := make(map[int]int)
    for i, num := range nums{
        if idx, ok := hashTable[target - num]; ok{
            return []int{idx,i}
        }
        hashTable[num] = i
    }
    return nil
} 