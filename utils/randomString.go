package utils

import (
	"crypto/rand"
	"log"
	"unsafe"
)

var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomString generates a random string of a length
func RandomString(len int) string {
	b := make([]byte, len)

	_, err := rand.Read(b)

	if err != nil {
		log.Printf("[ERROR] failed to random read: %+v", err)
		return ""
	}

	for i := 0; i < len; i++ {
		b[i] = alphabet[b[i]/5]
	}

	return *(*string)(unsafe.Pointer(&b))
}
