package regex

import "regexp"

var (
	ipRegex     = regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+$`)
	subnetRegex = regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+\/\d+$`)
	asnRegex    = regexp.MustCompile(`^\d+$`)
)

func IsValidIP(ip string) bool {
	return ipRegex.MatchString(ip)
}

func IsValidSubnet(subnet string) bool {
	return subnetRegex.MatchString(subnet)
}

func IsValidASN(asn string) bool {
	return asnRegex.MatchString(asn)
}
