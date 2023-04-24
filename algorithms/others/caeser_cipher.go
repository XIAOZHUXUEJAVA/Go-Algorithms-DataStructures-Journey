package others

func encrypt(plainText string, shiftKey int) string {
	bytes := []byte(plainText)
	for i, c := range bytes {
		if c >= 'a' && c <= 'z' {
			bytes[i] = 'a' + byte((int(c-'a')+shiftKey)%26)
		} else {
			bytes[i] = 'A' + byte((int(c-'A')+shiftKey)%26)
		}
	}
	return string(bytes)
}
func decrypt(cipherText string, shiftKey int) string {
	bytes := []byte(cipherText)
	for i, c := range bytes {
		if c >= 'a' && c <= 'z' {
			bytes[i] = 'a' + byte((int(c-'a')-shiftKey+26)%26)
		} else {
			bytes[i] = 'A' + byte((int(c-'A')-shiftKey+26)%26)
		}
	}
	return string(bytes)
}
