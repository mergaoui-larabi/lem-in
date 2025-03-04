package helpers

func Contains(str string, slice []string) bool {
	for _, s := range slice {
		if str == s {
			return true
		}
	}
	return false
}

func Split(b []byte) [][]byte {
	var buffer []byte
	var result [][]byte
	for i := range b {
		if b[i] == '\n' {
			result = append(result, buffer)
			buffer = []byte(nil)
			continue
		}
		buffer = append(buffer, b[i])
	}
	return result
}
