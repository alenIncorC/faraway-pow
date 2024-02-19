package verification

import (
	"bytes"
	"crypto/sha256"
	"errors"
)

var (
	ErrInvalidDifficulty = errors.New("invalid difficulty provided")
)

const (
	difficulty    = 3 // set difficulty level for POW challenge
	maxDifficulty = 10
)

type POWVerifier struct {
	difficulty int
}

func NewPOWPOWVerifier(difficulty int) (*POWVerifier, error) {
	if difficulty < 1 || difficulty > maxDifficulty {
		return nil, ErrInvalidDifficulty
	}

	return &POWVerifier{difficulty: difficulty}, nil
}

// VerifyPOW challenge
func (pv *POWVerifier) Verify(challenge, solution []byte) bool {
	if !bytes.HasPrefix(solution, challenge) {
		return false
	}

	hash := sha256.Sum256(solution)
	for i := 0; i < difficulty; i++ {
		if hash[i] > 0 {
			return false
		}
	}

	return true
}
