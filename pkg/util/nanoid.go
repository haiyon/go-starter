package util

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomNanoString random Nano String
func RandomNanoString(l int) string {
	id, _ := gonanoid.Generate(alphabet, l)
	return id
}

// RandomNanoID random Nano ID
func RandomNanoID() string {
	id, _ := gonanoid.Generate(alphabet, 22)
	return id
}
