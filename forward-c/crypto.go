package main

import (
	"king-go/crypto/classical"
)

var g_k3 classical.K3XsxSalt

func initCrypto(key string) {
	var sum byte = 0
	b := []byte(key)
	for i := 0; i < len(b); i++ {
		sum += b[i]
	}
	g_k3.Salt = sum

	g_k3.Encryption(b)
}
func Encryption(b []byte) ([]byte, error) {
	return g_k3.Encryption(b)
}
func Decryption(b []byte) ([]byte, error) {
	return g_k3.Decryption(b)
}
