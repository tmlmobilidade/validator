package lib

import (
	"crypto/rand"
)

const validationIDAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateValidationID creates a short id for correlating logs from one run.
func GenerateValidationID() string {
	const size = 5

	random := make([]byte, size)
	if _, err := rand.Read(random); err != nil {
		return "00000"
	}

	id := make([]byte, size)
	for i, b := range random {
		id[i] = validationIDAlphabet[int(b)%len(validationIDAlphabet)]
	}

	return string(id)
}
