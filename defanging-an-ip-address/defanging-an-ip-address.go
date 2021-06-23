func defangIPaddr(address string) string {
    result := strings.ReplaceAll(address,".","[.]")
    return result
}