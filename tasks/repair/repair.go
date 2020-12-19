package repair

import (
	"crypto/md5"
	hex "encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GetRepair() int {
	max := 99999999
	min := 10000000
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	number := random.Intn(max-min) + min
	return number
}

func CheckRepair(number int, hexResult string) bool {
	if number >= 100000000 && number < 10000000 {
		return false
	}
	a := md5.Sum([]byte(strconv.Itoa(number)))
	hexNumber := hex.EncodeToString(a[:])
	fmt.Println(hexNumber)
	if hexNumber == hexResult {
		return true
	}
	return false
}
