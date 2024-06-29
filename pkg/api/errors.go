package api

import "fmt"

type Error map[string]interface{}

func (e Error) ErrorCode() int {
	return int((e["code"]).(float64))
}

var ERRORS = map[int]string{
	100: "unknown error.",
	101: "missing api, version, or method parameter",
	102: "the requested API does not exist",
	103: "the requested method does not exist",
	104: "the requested version does not support the functionality",
	105: "the logged in session does not have permission",
	106: "session timeout",
	107: "session interrupted by duplicated login",
	108: "failed to upload file",
	109: "the network connection is unstable or the system is busy",
	110: "the network connection is unstable or the system is busy",
	111: "the network connection is unstable or the system is busy",
	114: "lost parameters for this API",
	115: "not allowed to upload a file",
	116: "not allowed to perform for a demo site.",
	117: "the network connection is unstable or the system is busy",
	118: "the network connection is unstable or the system is busy",
	119: "invalid session",
	150: "request source IP does not match the login IP",

	400: "No such account or incorrect password",
	401: "Disabled account",
	402: "Denied permission",
	403: "2-factor authentication code required",
	404: "Failed to authenticate 2-factor authentication code",
	406: "Enforce to authenticate with 2-factor authentication code",
	407: "Blocked IP source",
	408: "Expired password cannot change",
	409: "Expired password",
	410: "Password must be changed",
}

func (e Error) ErrorCodeDescription() string {
	s, ok := ERRORS[e.ErrorCode()]
	if ok {
		return fmt.Sprintf("%d: %s. ", e.ErrorCode(), s)
	}
	return fmt.Sprintf("unfamiliar error code: %d. ", e.ErrorCode())
}
