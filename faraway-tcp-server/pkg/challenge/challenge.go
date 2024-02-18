package challenge

import (
	"math/rand"

	"faraway-tcp-server/pkg/server/tcpserver/protocol"
)

func GetRandomChallengeMessage() []byte {
	rndMSG := append(protocol.ChallengeType, GenerateRandomChallenge()...)
	return rndMSG
}

func GenerateRandomChallenge() []byte {
	challenges := []string{
		"FARAWAY\n",
		"Mini Royal Nations\n",
		"MiniNations\n",
		"Web3 Gaming\n",
	}
	quoteIndex := rand.Intn(len(challenges))
	return []byte(challenges[quoteIndex])
}
