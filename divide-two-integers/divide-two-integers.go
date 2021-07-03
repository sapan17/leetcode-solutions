func divide(dividend int, divisor int) int {
    if (dividend/divisor) > 0{
        if (dividend/divisor) < 2147483648{
            return (dividend/divisor)
        } else{
            return 2147483647
        }
    } else{
        if (dividend/divisor) < -2147483648{
            return -2147483648
        } else{
            return (dividend/divisor)
        }
    }
}