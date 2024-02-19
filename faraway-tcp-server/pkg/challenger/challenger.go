package challenger

type Challenger interface {
	GetChallenge() []byte
	GetDifficulty() ([]byte, error)
}
