package usecase

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"net"
	"os"

	"github.com/alenIncorC/faraway-test/faraway-tcp-client/config"
	"github.com/alenIncorC/faraway-test/faraway-tcp-client/pkg/solver"
	"github.com/alenIncorC/faraway-test/libs/protocol"
	"github.com/alenIncorC/faraway-test/libs/utils"
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
			fmt.Errorf("failed to get quote: %s\n", err.Error())
		} else {
			fmt.Println(string(q))
		}
	}

	return nil
}

// GetQuote returns a quote from the server
func (c *Client) GetQuote(ctx context.Context) ([]byte, error) {
	var dialer net.Dialer
	conn, err := dialer.DialContext(ctx, "tcp", c.conf.Addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
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
	challenge, err := utils.ReadMessage(conn)
	if err != nil {
		return nil, fmt.Errorf("receive challenge err: %w", err)
	}

	// send solution
	solution := c.solver.Solve(challenge)
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

func HandleConnection(conn net.Conn) {
	challenge := new(entity.Challenge)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		serverResponse := scanner.Bytes()
		HandleResponse(conn, serverResponse, challenge)
	}
}

func HandleResponse(conn net.Conn, serverResponse []byte, challenge *entity.Challenge) {
	serverResponse = bytes.Trim(serverResponse, "\n")
	if bytes.Equal(serverResponse, []byte("STOP")) {
		fmt.Println("Server hang up")
		conn.Close()
		os.Exit(0)
	}

	msgT, msgV := GetMessage(serverResponse)

	//custom tcp message protocol
	switch msgT {
	case protocol.Challenge:
		challenge.Challenge = msgV
	case protocol.Difficulty:
		challenge.Difficulty = msgV
	case protocol.Wisdom:
		fmt.Println(string(msgV))
		challenge = &entity.Challenge{}
		conn.Write([]byte("STOP\n"))
		conn.Close()
	case protocol.Message:
		fmt.Println(string(msgV))
		conn.Write([]byte("STOP\n"))
	default:
		fmt.Println("Unknown server response: " + string(msgV))
		conn.Close()
		os.Exit(1)
	}

	if challenge.Challenge != nil && challenge.Difficulty != nil {
		result := miner.Mine(challenge.Challenge, challenge.Difficulty)
		conn.Write(result)
		conn.Write([]byte("\n"))
	}
}

func GetMessage(data []byte) (int, []byte) {
	allTypes := map[int][]byte{
		protocol.Challenge:  protocol.ChallengeType,
		protocol.Difficulty: protocol.DifficultyType,
		protocol.Wisdom:     protocol.WisdomType,
		protocol.Message:    protocol.MessageType,
	}

	for k, v := range allTypes {
		if bytes.HasPrefix(data, v) {
			return k, data[len(v):]
		}
	}

	return protocol.Unknown, nil
}
