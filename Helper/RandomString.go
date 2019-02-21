package Helper

import "math/rand"

func String(length int) string {
// initialize set params
letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// randomize string
b := make([]byte, length)
for i := range b {
b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
}

return string(b)
}
