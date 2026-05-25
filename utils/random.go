package utils

import (
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {

	rand.Seed(time.Now().UnixNano())

	result := make([]byte, length)

	for i := range result {

		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
