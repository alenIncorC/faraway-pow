package protocol

import (
	"fmt"
)

const Prefix string = "TypeCode"

const (
	Unknown = iota
	Challenge
	Difficulty
	Wisdom
	Message
)

var (
	ChallengeType  []byte = []byte(fmt.Sprintf("%s%d: ", Prefix, Challenge))
	DifficultyType []byte = []byte(fmt.Sprintf("%s%d: ", Prefix, Difficulty))
	WisdomType     []byte = []byte(fmt.Sprintf("%s%d: ", Prefix, Wisdom))
	MessageType    []byte = []byte(fmt.Sprintf("%s%d: ", Prefix, Message))
)
