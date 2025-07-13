package logic

import (
	"log"
	"time"
)

func RefreshTokens() {

	for {
		log.Println("refreshing users with tokens")

		usersIps := GetAllUsers()
		for _, ip := range usersIps {
			FillToken(ip)
			PurgeOldUsers(ip)
		}

		time.Sleep(time.Second * 5)
	}
}
