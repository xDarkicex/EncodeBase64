package encoder

import "strings"

var b64 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
var base = len(b64)

func IntToBase64(number int) string {
	str := ""
	for number > 0 {
		r := number % base
		number -= r
		number /= base
		str = string(b64[r]) + str
	}
	return str
}

func Base64ToInt(base64 string) int {
	number := 0
	sliced := strings.Split(base64, "")
	for _, v := range sliced {
		r := strings.Index(b64, v)
		number *= base
		number += r
	}
	return number
}
