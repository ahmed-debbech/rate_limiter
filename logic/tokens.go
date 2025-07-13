package logic

import (
	"errors"
	"log"
)

func ConsumeToken(originalIp string) error {

	log.Println("consuming tokens")
	bucket := GetBucket(originalIp)

	if bucket.TokensLeft == 0 {
		return errors.New("Not enough tokens for " + originalIp)
	}
	bucket.TokensLeft--
	log.Println("tokens left:", bucket.TokensLeft)

	return nil
}

func FillToken(originalIP string) {

	bucket := GetBucket(originalIP)
	if bucket.TokensLeft == bucket.TokenMaxLimit {
		return
	}

	bucket.TokensLeft++
}
