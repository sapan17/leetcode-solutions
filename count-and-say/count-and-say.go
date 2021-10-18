import (
    "strconv"
    "strings"
)
        
func countAndSay(n int) string {
    result := "1"
    for i := 0; i <= n-2; i++ {
        resultTokens := make([]string, 20)
        for j := 0; j < len(result);{
            num := result[j]
            count := 0
            for j < len(result) && num == result[j] {
                count++
                j++
            }
            resultTokens = append(resultTokens, strconv.Itoa(count), string(num))
        }
        result = strings.Join(resultTokens, "")
    }
    return result
}