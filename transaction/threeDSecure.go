package transaction

type ThreeDSecureEnrollmentStatus string

func IsThreeDSecureEnrollmentStatus(status string) bool {
	return threeDSecureEnrollmentStatus[status]
}

var threeDSecureEnrollmentStatus = map[string]bool{
	"Y": true,
	"B": true,
	"N": true,
	"U": true,
}

type ThreeDSecureAuthenticationStatus string

func IsThreeDSecureAuthenticationStatus(status string) bool {
	return threeDSecureAuthenticationStatus[status]
}

var threeDSecureAuthenticationStatus = map[string]bool{
	"Y": true,
	"N": true,
	"A": true,
	"C": true,
	"R": true,
	"U": true,
}

type ThreeDSecureSignatureVerification string

func IsThreeDSecureSignatureVerification(value string) bool {
	return threeDSecureSignatureVerification[value]
}

var threeDSecureSignatureVerification = map[string]bool{
	"Y": true,
	"N": true,
}

// type ThreeDSecureErrorNo 	string

// func IsThreeDSecureErrorNo(value string) bool {
// 	if _, err := strconv.Atoi(value); err != nil {
// 		return false
// 	}
// 	return true
// }
