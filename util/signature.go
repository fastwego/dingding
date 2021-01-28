package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

func Signature(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(fmt.Sprintf("%v\n%s", time.Now().UnixNano()/1e6, data)))
	return url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
}
