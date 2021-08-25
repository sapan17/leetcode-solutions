func countPrimes(n int) int {
    if n<3{
        return 0
    }
    
    primecount := 0
    nums := make([]bool,n)
    
    for i:=0; i<n; i++{
        nums[i] = true
    }   
    for i :=2; i*i <n; i++{
        if nums[i] == true{
            for j:=2; j*i<n;j++{
                nums[j*i] = false
            }
        }
    }
    for i:=2; i<n; i++{
        if nums[i] == true{
            primecount +=1
        }
    }
    return primecount
}