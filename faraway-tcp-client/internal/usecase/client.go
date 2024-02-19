package usecase

import (
	"context"
	"fmt"
	"net"

	"github.com/alenIncorC/faraway-pow/faraway-tcp-client/config"
	"github.com/alenIncorC/faraway-pow/faraway-tcp-client/pkg/solver"
	"github.com/alenIncorC/faraway-pow/libs/utils"
)

// client represents a client
type Client struct {
	conf   *config.Config
	solver solver.Solver
}

// NewClient creates a new client
func NewClient(conf *config.Config, solver solver.Solver) *Client {
	return &Client{
		conf:   conf,
		solver: solver,
	}
}

// Start started fetch
func (c *Client) Start(ctx context.Context, count int) error {
	for i := 0; i < count; i++ {
		if ctx.Err() != nil {
			break
		}

		q, err := c.GetQuote(ctx)
		if err != nil {
			return fmt.Errorf("failed to get quote: %s\n", err.Error())
		} else {
			fmt.Println(string(q))
		}
	}

	return nil
}

// GetQuote returns a quote from the server
func (c *Client) GetQuote(ctx context.Context) ([]byte, error) {
	var dialer net.Dialer
	conn, err := dialer.DialContext(ctx, "tcp", c.conf.ServerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %s", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Errorf("failed to close connection %s\n", err.Error())
		}
	}()

	// challenge request
	if err := utils.WriteMessage(conn, []byte("challenge")); err != nil {
		return nil, fmt.Errorf("send challenge request err: %w", err)
	}

	// receive challenge
	puzzle, err := utils.ReadMessage(conn)
	if err != nil {
		return nil, fmt.Errorf("receive challenge err: %s", err)
	}

	difficulty, err := utils.ReadInt64(conn)
	if err != nil {
		return nil, fmt.Errorf("receive difficulty err: %s", err)
	}

	// send solution
	solution := c.solver.Solve(puzzle, difficulty)
	if err := utils.WriteMessage(conn, solution); err != nil {
		return nil, fmt.Errorf("send solution err: %w", err)
	}

	// receive quote
	quote, err := utils.ReadMessage(conn)
	if err != nil {
		return nil, fmt.Errorf("receive quote err: %w", err)
	}

	return quote, nil
}
