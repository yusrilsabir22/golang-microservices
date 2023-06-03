package utils

import "regexp"

// Pkg Validator should be better than this
func IsValidMSISDN(msisdn string) bool {
	re := regexp.MustCompile(`^\\{0,1}0{0,1}62[0-9]+$`)
	submatch := re.FindStringSubmatch(msisdn)
	return len(submatch) > 2
}
