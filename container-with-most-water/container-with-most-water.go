func maxArea(height []int) int {
    maxA := 0
    l, r := 0, len(height)-1
    for l < r{
        area := min(height[l],height[r])*(r-l)
        
        if area > maxA{
            maxA = area
        }
        
        if height[l] > height[r]{
            r--
        } else{
            l++
        } 
    }
    return maxA
}

func min(x int,y int)int{
    if x < y{
        return x
    }
    return y
}