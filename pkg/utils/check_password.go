package utils

func CheckPassword(password string) bool {
	if len(password) >= 6 {
		return true
	} else {
		return false
	}

}
