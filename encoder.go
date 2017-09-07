package encoder

func StringToBase64(number int) string {
	b64 := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
	base := len(b64)
	str := ""
	for number > 0 {
		r := number % base
		number -= r
		number /= base
		str = string(b64[r]) + str
	}
	return str
}

func Base64ToString(base64 string) int {
	return 0
}
