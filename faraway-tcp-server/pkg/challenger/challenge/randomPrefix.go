package challenge

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"

	"github.com/alenIncorC/faraway-pow/libs/random"
)

type Challenge struct {
	difficulty int
}

func NewChallenge(difficulty int) (*Challenge, error) {
	if difficulty < 1 || difficulty > 6 {
		return nil, fmt.Errorf("invalid difficulty")
	}

	return &Challenge{difficulty: difficulty}, nil
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

func (c *Challenge) GetDifficulty() ([]byte, error) {
	buf := new(bytes.Buffer)
	num := uint16(c.difficulty)
	err := binary.Write(buf, binary.LittleEndian, num)
	if err != nil {
		return nil, fmt.Errorf("binary.Write failed: %s", err)
	}

	return buf.Bytes(), nil
}
