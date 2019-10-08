package transaction

import (
	"regexp"
)

const regexV4 = `\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}\b`

func IsIPv4(ipv4 string) bool {
	re := regexp.MustCompile(regexV4)
	return re.MatchString(ipv4)
}
