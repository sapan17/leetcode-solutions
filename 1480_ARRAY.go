func runningSum(nums []int) []int {
	presum := 0 // presum variable is first initialized
	for i := 0; i < len(nums); i++ {
		presum = presum + nums[i] // In the first loop, the value of presum is 0, and the first value of nums is added to this variable
		nums[i] = presum          // Later the nums[i] value is passed to the presum variable
	}
	return nums
}

// SOLUTION WITH EXAMPLE
// INPUT: [1,2,3,4]
// OUTPUT: [1,3,6,10]

// presum = presum + nums[i]
// 1st loop
// presum = 0 + 1 = 1
// 2nd loop
// presum = 1 + 2 = 3

// nums[i] = presum
// 1st loop
// nums[0] = 1
// 2nd loop
// nums[1] = 3

// While in the upper part, the temporary variable presum is getting added as per the array, in the below part the value in nums array is getting replaced by the temporary variable.