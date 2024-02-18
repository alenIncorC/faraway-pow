package solver

type Solver interface {
	Solve(challenge []byte) []byte
}
