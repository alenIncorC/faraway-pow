package challenge

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alenIncorC/faraway-pow/libs/random"
)

type Challenge struct {
	difficulty int64
}

func NewChallenge(difficulty int) (*Challenge, error) {
	if difficulty < 1 || difficulty > 6 {
		return nil, fmt.Errorf("invalid difficulty")
	}

	return &Challenge{difficulty: int64(difficulty)}, nil
}

func (c *Challenge) GetChallenge() []byte {
	return randomPrefix()
}

func randomPrefix() []byte {
	prefix := make([]byte, rand.Intn(20-7)+7)
	seed := uint64(time.Now().Unix())

	random.RandomString(prefix, 0, seed)

	return prefix
}

func (c *Challenge) GetDifficulty() int64 {
	return c.difficulty
}
