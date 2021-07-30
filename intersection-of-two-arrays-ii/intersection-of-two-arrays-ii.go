func intersect(nums1 []int, nums2 []int) []int {
    p1 := 0
    p2 := 0
    res := make([]int,0)
    cache := make(map[int]struct{},0)
    
    for{
        if p1 >= len(nums1){
            break
        }
        if p2 >= len(nums2){
            p2 = 0
            p1++
            continue
        }
        if nums1[p1] != nums2[p2]{
            p2++
            continue
        }
        if nums1[p1] == nums2[p2]{
            _,ok := cache[p2]
            if ok{
                p2++
                continue
            } else{
                cache[p2] = struct {}{}
                res = append(res,nums1[p1])
                p2 = 0
                p1++
            }
        }
    }
    return res
}