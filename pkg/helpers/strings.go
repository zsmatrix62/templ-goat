package helpers

import "golang.org/x/exp/rand"

func RandomString(len int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, len)
	for i := range b {
		b[i] = chars[rand.Intn(len)]
	}
	return string(b)
}
