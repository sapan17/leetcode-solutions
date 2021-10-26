func productExceptSelf(nums []int) []int {
    n := len(nums)
    
    arr1 := make([]int, n)
    arr2 := make([]int, n)
    arr3 := make([]int, n)
    
    arr1[0] = 1
    arr2[n-1] = 1
    
    for i:=1; i<n; i++{
        arr1[i] = arr1[i-1] * nums[i-1]
    }
    for i:= n-2; i>=0; i--{
        arr2[i] = arr2[i+1] * nums[i+1]
    }
    for i:=0; i<n; i++{
        arr3[i] = arr1[i] * arr2[i]
    }
    return arr3
}