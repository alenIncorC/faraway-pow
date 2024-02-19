package verifier

type Verifier interface {
	Verify(challenge, solution []byte) bool
}
