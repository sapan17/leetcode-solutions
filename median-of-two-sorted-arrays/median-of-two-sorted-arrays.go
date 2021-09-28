func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var nums3 []int
    for i:=0; i<len(nums1); i++{
        nums3 = append(nums3,nums1[i])
    }
    for i:=0; i<len(nums2); i++{
        nums3 = append(nums3,nums2[i])
    }
    //nums1 = append(nums1,nums2...)
    sort.Ints(nums3)
    
    if len(nums3)%2 == 0{
        return float64(nums3[len(nums3)/2 - 1] + nums3[len(nums3)/2])/2
    } else{
        return float64(nums3[len(nums3)/2])
    }
}