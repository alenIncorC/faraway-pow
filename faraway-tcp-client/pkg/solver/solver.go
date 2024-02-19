package solver

type Solver interface {
	Solve(challenge []byte, difficulty int64) []byte
}
