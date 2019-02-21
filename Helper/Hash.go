package Helper

import (
	"crypto/sha256"
	"encoding/base64"
)

func Make(feed string) string {
	sha256Crypter := sha256.New()
	sha256Crypter.Write([]byte(feed))

	return base64.URLEncoding.EncodeToString((sha256Crypter.Sum(nil)))
}

func Check(feed string, challenge string) bool {
	hashedFeed := Make(feed)

	if hashedFeed != challenge {
		return false
	}

	return true
}