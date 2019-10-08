package transaction

import "regexp"

type IPv4 string

const regexV4 = `\b(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.
	(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.
	(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.
	(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\b`

func IsIPv4(ipv4 string) bool {
	re := regexp.MustCompile(regexV4)

	return re.MatchString(ipv4)
}
