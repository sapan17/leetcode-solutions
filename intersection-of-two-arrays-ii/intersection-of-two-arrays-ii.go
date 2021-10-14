func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var (
		i,j int
		ret = []int{}
	)

	for i, j = 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			ret = append(ret, nums1[i])
			i++
			j++
        } else if nums1[i] < nums2[j] {
            i++
        } else{
            j++
        }
	}
    
	return ret
}