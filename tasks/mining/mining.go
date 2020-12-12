package mining

import (
	md5 "crypto/md5"
	"encoding/hex"
	"strings"
)

func CheckMining(key string, text string) bool {
	hash := ConvertToHash(text)
	if len(strings.Split(hash, key)) >= 3 {
		return true
	}
	return false
}

func ConvertToHash(text string) string {
	md5Hash := md5.Sum([]byte(text))
	return hex.EncodeToString(md5Hash[:])
}
