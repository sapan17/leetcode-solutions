func maxArea(height []int) int {
    A := 0
    l := 0
    r := len(height) - 1
    mini := height[0]
    maxA := 0
    for l < r{
        mini = min(height[l], height[r])
        A = mini * (r-l)
        if maxA < A{
            maxA = A
        }
        if height[l] > height[r]{
            r -= 1
        } else{
            l += 1
        }
    }
    return maxA
}

func min(a,b int) int{
    if a > b{
        return b
    } else{
        return a
    }
}