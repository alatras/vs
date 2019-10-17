package transaction

import (
	"regexp"
)

const regexV4 = `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`

var regex = regexp.MustCompile(regexV4)

func IsIPv4(ipv4 string) bool {
	return regex.MatchString(ipv4)
}
