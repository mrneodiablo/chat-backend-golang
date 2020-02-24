package utilities

import (
	"math/rand"
	"time"
)


func GenerateMessageId(userId int64) int64 {

	timeTmp := time.Now().UnixNano() / int64(time.Millisecond)
	data := timeTmp + userId + int64(Random())
	return data


}

func Random() int {
	max := 10000
	min := 10
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + max
}
