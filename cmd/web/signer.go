package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	goalone "github.com/bwmarrin/go-alone"
)

var secret = os.Getenv("SIGNER_SECRET")

var secretKey []byte

// NewURLSigner creates a new signer
func NewURLSigner() {
	secretKey = []byte(secret)
}

// GenerateTokenFromString generates a signed token
func GenerateTokenFromString(data string) string {
	var urlToSign string

	s := goalone.New(secretKey, goalone.Timestamp)
	if strings.Contains(data, "?") {
		urlToSign = fmt.Sprintf("%s&hash=", data)
	} else {
		urlToSign = fmt.Sprintf("%s?hash=", data)
	}

	tokenBytes := s.Sign([]byte(urlToSign))
	token := string(tokenBytes)

	return token
}

// VerifyToken verifies a signed token
func VerifyToken(token string) bool {
	s := goalone.New(secretKey, goalone.Timestamp)
	_, err := s.Unsign([]byte(token))

	// false if signature is not valid. Token was tampered with, forged
	return err == nil
}

// Expired checks to see if a token has expired
func Expired(token string, minutesUntilExpire int) bool {
	s := goalone.New(secretKey, goalone.Timestamp)
	ts := s.Parse([]byte(token))

	// time.Duration(seconds)*time.Second
	return time.Since(ts.Timestamp) > time.Duration(minutesUntilExpire)*time.Minute
}
