package password

func PasswordVerify(password string, hash string) bool {
	errw := CompareHashAndPassword([]byte(hash), []byte(password))
	if errw != nil {
		return false
	}
	return true
}

func PasswordHash(password string) string {
	r1, r2 := GenerateFromPassword([]byte(password), 10)
	if r2 != nil {
		return ""
	}
	return string(r1)
}
