package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"time"
)

// implements DNS Made Easy's HMAC function via https://api-docs.dnsmadeeasy.com/#5b98221f-37e9-4845-a349-5e959241b4a5
func generateHMAC(secretKey string) (string, string) {
	timeLocation, _ := time.LoadLocation("GMT")
	currentTime := time.Now().In(timeLocation).Format(time.RFC1123)
	crypto := hmac.New(sha1.New, []byte(secretKey))
	crypto.Write([]byte(currentTime))
	hmacResult := hex.EncodeToString(crypto.Sum(nil))
	return hmacResult, currentTime
}
