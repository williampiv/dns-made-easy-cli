package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"time"
)

// implements DNS Made Easy's HMAC function via WEBADDRESS
func generateHMAC(secretKey string) (string, string) {
	timeLocation, _ := time.LoadLocation("GMT")
	currentTime := time.Now().In(timeLocation).Format(time.RFC1123)
	crypto := hmac.New(sha1.New, []byte(secretKey))
	crypto.Write([]byte(currentTime))
	hmacResult := hex.EncodeToString(crypto.Sum(nil))
	return hmacResult, currentTime
}
