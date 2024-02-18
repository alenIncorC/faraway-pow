package verification

import (
	"crypto/sha256"
	"strconv"

	"faraway-tcp-server/pkg/server/tcpserver/protocol"
)

const difficulty = 3 // set difficulty level for POW challenge

func GetDifficultyMessage() []byte {
	bytes := append(protocol.DifficultyType, []byte(strconv.Itoa(difficulty))...)
	bytes = append(bytes, []byte("\n")...)
	return bytes
}

// VerifyPOW challenge
func VerifyPOW(challenge []byte) bool {
	hash := sha256.Sum256(challenge)
	for i := 0; i < difficulty; i++ {
		if hash[i] > 0 {
			return false
		}
	}
	return true
}
