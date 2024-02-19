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

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) Solve(challenge []byte, difficulty int64) []byte {
	start := time.Now()

	numberOfGOR := runtime.NumCPU()
	hashes := make([]int, numberOfGOR)
	solutionChan := make(chan []byte)
	done := make(chan struct{})

	for i := 0; i < numberOfGOR; i++ {
		go func(index int, cmplx int64) {
			offset := len(challenge)
			strWithPrefix := make([]byte, 20+offset)
			copy(strWithPrefix[:offset], challenge)
			seed := uint64(index)

			for {
				select {
				case <-solutionChan:
					fmt.Printf("goroutine num %d - I am done\n", index)
					return
				default:
					hashes[index]++
					seed = random.RandomString(strWithPrefix, offset, seed)
					if hash(strWithPrefix, cmplx, done) {
						fmt.Printf("goroutine num %d - found soltion %s\n", index, strWithPrefix)
						solutionChan <- strWithPrefix
						return
					}
				}
			}
		}(i, difficulty)
	}

	solution := <-solutionChan
	close(solutionChan)
	done <- struct{}{}

	hashesSum := 0

	for i := 0; i < numberOfGOR; i++ {
		fmt.Printf("goroutine num: %d,number of hashes %d\n", i, hashes[i])
		hashesSum += hashes[i]
	}

	end := time.Now()
	fmt.Printf("%s - totalNumber of hashes: %d, time spent: %g, processed: %s, hashesh per sec: %s\n",
		solution, hashesSum, end.Sub(start).Seconds(), humanize.Comma(int64(hashesSum)), humanize.Comma(int64(float64(hashesSum)/end.Sub(start).Seconds())))

	return solution
}

func hash(str []byte, complexity int64, done <-chan struct{}) bool {
	hash := sha256.Sum256(str)
	for i := 0; i < int(complexity); i++ {
		select {
		case <-done:
			return false
		default:
			if hash[i] > 0 {
				return false
			}
		}
	}

	return true
}
