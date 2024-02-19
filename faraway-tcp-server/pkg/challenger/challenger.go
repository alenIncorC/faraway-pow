package challenger

type Challenger interface {
	GetChallenge() []byte
	GetDifficulty() int64
}
