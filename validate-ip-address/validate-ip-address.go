import "regexp"

const i8 string = `([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
var v4 = regexp.MustCompile("^" + i8 + `(\.` + i8 + `){3}$`)

const hex string = "[0-9a-f]{1,4}"
var v6 = regexp.MustCompile("^(?i)[1-9a-f][0-9a-f]{0,3}(:" + hex + "){7}$")

func validIPAddress(IP string) string {
    if v4.Match([]byte(IP)) {
        return "IPv4"
    }
    
    if v6.Match([]byte(IP)) {
        return "IPv6"
    }
    
    return "Neither"
}