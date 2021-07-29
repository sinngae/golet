package debug_

func IsDebugging(debug bool) bool {
	if Debug {
		return true
	}
	if debug {
		return true
	}
	return false
}
