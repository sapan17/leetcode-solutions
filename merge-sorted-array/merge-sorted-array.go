func merge(nums1 []int, m int, nums2 []int, n int)  {
    i := len(nums1)-1
    for n>0{
        if m==0 || nums1[m-1] <= nums2[n-1]{
            nums1[i]=nums2[n-1]
            i--
            n--
        } else{
            nums1[i]=nums1[m-1]
            i--
            m--
        }
    }
}