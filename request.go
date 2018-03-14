package acm

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"time"
)

const (
	headerAccessKey = "Spas-AccessKey"
	headerTS        = "timeStamp"
	headerSignature = "Spas-Signature"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func getSign(encryptText, encryptKey string) string {
	key := []byte(encryptKey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(encryptText))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
