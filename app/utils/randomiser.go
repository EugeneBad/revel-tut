package utils


import (
"time"
"math/rand"
)

func Randomiser() string{
	rand.Seed(time.Now().UnixNano())

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	randBytes := []byte{}

	for i := 0; i < 15; i++ {
		randBytes = append(randBytes, letterBytes[rand.Intn(52)])
	}
	return string(randBytes)
}