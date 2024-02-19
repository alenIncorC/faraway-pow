package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/alenIncorC/faraway-pow/faraway-tcp-server/config"
	"github.com/alenIncorC/faraway-pow/faraway-tcp-server/internal/repository"
	"github.com/alenIncorC/faraway-pow/faraway-tcp-server/pkg/challenger"
	"github.com/alenIncorC/faraway-pow/faraway-tcp-server/pkg/verifier"
	"github.com/alenIncorC/faraway-pow/libs/utils"
)

// server represents a server
type server struct {
	cfg        *config.Config
	verifier   verifier.Verifier
	challenger challenger.Challenger
	repo       repository.Repositories
	listener   net.Listener
	wg         sync.WaitGroup
	cancel     context.CancelFunc
}

// NewServer creates a new server
func NewServer(
	cfg *config.Config,
	verifier verifier.Verifier,
	challenger challenger.Challenger,
	repo repository.Repositories,
) *server {
	return &server{
		cfg:        cfg,
		verifier:   verifier,
		challenger: challenger,
		repo:       repo,
	}
}

// Run starts the server
func (s *server) Run(ctx context.Context) (err error) {
	ctx, s.cancel = context.WithCancel(ctx)
	defer s.cancel()

	lc := net.ListenConfig{
		KeepAlive: s.cfg.KeepAlive,
	}
	s.listener, err = lc.Listen(ctx, "tcp", s.cfg.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	log.Printf("server started on port %s\n", s.listener.Addr().String())

	s.wg.Add(1)
	go s.serve(ctx)
	s.wg.Wait()

	log.Println("server stopped")

	return nil
}

// Stop stops the server
func (s *server) Stop() {
	s.cancel()
}

func (s *server) serve(ctx context.Context) {
	defer s.wg.Done()

	go func() {
		<-ctx.Done()
		err := s.listener.Close()
		if err != nil && !errors.Is(err, net.ErrClosed) {
			fmt.Println("failed to close listener: ", err.Error())
		}
	}()

	for {
		conn, err := s.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			log.Println("listener closed")
			return
		} else if err != nil {
			log.Println("failed to accept connection: ", err.Error())
			continue
		}

		s.wg.Add(1)
		go func(conn net.Conn) {
			defer s.wg.Done()

			if err = s.handle(conn); err != nil {
				fmt.Println("handle error: ", err.Error())
			}
		}(conn)
	}
}

func (s *server) handle(conn net.Conn) error {
	defer conn.Close()
	_ = conn.SetDeadline(time.Now().Add(s.cfg.Deadline))

	// receive challenge request
	if _, err := utils.ReadMessage(conn); err != nil {
		return fmt.Errorf("read message err: %s", err)
	}

	// send challenge
	challenge := s.challenger.GetChallenge()

	if err := utils.WriteMessage(conn, challenge); err != nil {
		return fmt.Errorf("send challenge err: %s", err)
	}

	//send difficulty
	difficulty := s.challenger.GetDifficulty()

	if err := utils.WriteInt64(conn, difficulty); err != nil {
		return fmt.Errorf("send challenge err: %w", err)
	}

	// receive solution
	solution, err := utils.ReadMessage(conn)
	if err != nil {
		return fmt.Errorf("receive proof err: %s", err)
	}

	// verify solution
	if !s.verifier.Verify(challenge, solution) {
		return fmt.Errorf("solution not accepted")
	}

	// send result
	quote, err := s.repo.Quotes.GetQuote()
	if err != nil {
		return fmt.Errorf("get quote err: %w", err)
	}

	if err = utils.WriteMessage(conn, []byte(quote)); err != nil {
		return fmt.Errorf("send quote err: %w", err)
	}

	return nil
}
