func peakIndexInMountainArray(arr []int) int {
    n := len(arr)
    l := 0
    r := n -1
    
    for l < r{
        mid := l + (r-l) / 2
        if arr[mid] < arr[mid+1]{
            l = mid + 1
        } else{
            r = mid
        }
    }
    return l
}
    
