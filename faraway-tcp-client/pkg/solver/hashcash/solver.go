package hashcash

import (
	"crypto/sha256"
	"fmt"
	"runtime"
	"time"

	"github.com/alenIncorC/faraway-pow/libs/random"
	humanize "github.com/dustin/go-humanize"
)

var (
	ErrInvalidDifficulty = fmt.Errorf("invalid difficulty")
)

type Solver struct {
	Difficulty int
}

func NewSolver(difficulty int) (*Solver, error) {
	if difficulty < 1 || difficulty > 8 {
		return nil, ErrInvalidDifficulty
	}

	return &Solver{Difficulty: difficulty}, nil
}

func (s *Solver) Solve(challenge []byte) []byte {
	start := time.Now()
	complexity := s.Difficulty

	numberOfGOR := runtime.NumCPU()
	hashes := make([]int, numberOfGOR)
	solutionChan := make(chan []byte)
	done := make(chan struct{})

	for i := 0; i < numberOfGOR; i++ {
		go func(index int, cmplx int) {
			defer close(solutionChan)

			offset := len(challenge)
			strWithPrefix := make([]byte, 20+offset)
			copy(strWithPrefix[:offset], challenge)
			seed := uint64(index)

			for {
				select {
				case <-done:
					return
				default:
					hashes[index]++
					seed = random.RandomString(strWithPrefix, offset, seed)
					if hash(strWithPrefix, cmplx) {
						done <- struct{}{}
						solutionChan <- strWithPrefix
						break
					}
				}
			}
		}(i, complexity)
	}

	solution := <-solutionChan

	hashesSum := 0

	for i := 0; i < numberOfGOR; i++ {
		fmt.Printf("goroutine num: %d,number of hashes %d\n", i, hashes[i])
		hashesSum += hashes[i]
	}

	end := time.Now()
	fmt.Println(string(solution))
	fmt.Printf("totalNumber of hashes: %d\n", hashesSum)
	fmt.Printf("time spent: %g\n", end.Sub(start).Seconds())
	fmt.Println("processed:", humanize.Comma(int64(hashesSum)))
	fmt.Println("hashesh per sec:", humanize.Comma(int64(float64(hashesSum)/end.Sub(start).Seconds())))
	return solution
}

func hash(str []byte, complexity int) bool {
	hash := sha256.Sum256(str)
	for i := 0; i < complexity; i++ {
		if hash[i] > 0 {
			return false
		}
	}

	return true
}
