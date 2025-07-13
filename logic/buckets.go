package logic

import (
	"log"
	"time"
)

const MAX_TOKENS_PER_USER = 10

type Bucket struct {
	TokensLeft    int
	TokenMaxLimit int
	LastVisit     int
}

var userMap map[string]*Bucket = make(map[string]*Bucket, 0)

func GetBucket(originalIp string) *Bucket {
	bucket, ok := (userMap[originalIp])
	if !ok {
		// add new user with new bucket

		buc := &Bucket{
			TokensLeft:    MAX_TOKENS_PER_USER,
			TokenMaxLimit: MAX_TOKENS_PER_USER,
			LastVisit:     int(time.Now().Unix()),
		}
		userMap[originalIp] = buc
		return buc
	}

	return bucket

}

func GetAllUsers() []string {

	log.Println(userMap)
	users := make([]string, 0)
	for key, _ := range userMap {
		users = append(users, key)
	}

	return users
}

func PurgeOldUsers(originalIp string) {

	if (int(time.Now().Unix()) - userMap[originalIp].LastVisit) > 60 { //60 seconds
		delete(userMap, originalIp)
	}
}
