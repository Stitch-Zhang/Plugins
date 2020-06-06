package randomChar

import (
	"math/rand"
	"time"
)

func Creat(lent int) (result string) {
	var Char string
	for i := 0; i < lent; i++ {
		rand.Seed(time.Now().Unix() + int64(100*i))
		switch rand.Intn(3) {
		case 0:
			Char = string(65 + rand.Intn(25))
		case 1:
			Char = string(97 + rand.Intn(25))
		case 2:
			Char = string(48 + rand.Intn(9))
		}
		result = result + Char
	}
	return
}
