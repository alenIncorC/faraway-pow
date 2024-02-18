package pow

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"

	"faraway-tcp-server/pkg/challenge"
	"faraway-tcp-server/pkg/quotes"
	"faraway-tcp-server/pkg/server/tcpserver/protocol"
	"faraway-tcp-server/pkg/verification"
)

// handle incoming connections
func HandleConnection(conn net.Conn) {
	defer conn.Close()

	challenge := challenge.GetRandomChallengeMessage()
	difficulty := verification.GetDifficultyMessage()
	//write challenge and difficulty
	conn.Write(challenge)
	conn.Write(difficulty)

	for {
		// read POW challenge from client
		reader := bufio.NewReader(conn)
		clientResponse, err := reader.ReadBytes(byte('\n'))
		if err == io.EOF {
			fmt.Println("connection with client closed")
			conn.Close()
			return
		}

		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}

		clientResponse = bytes.Trim(clientResponse, "\n")

		if string(clientResponse) == "STOP" {
			conn.Write([]byte("Stop connection\n"))
			return
		}

		// verify POW challenge
		if !verification.VerifyPOW(clientResponse) {
			invalid := append(protocol.MessageType, []byte("Invalid POW challenge")...)
			invalid = append(invalid, []byte("\n")...)
			conn.Write(invalid)
			return
		}

		// generate random quote and send to client
		quote := quotes.GetQuote()
		conn.Write(append(protocol.WisdomType, []byte(quote+"\n")...))
		conn.Write([]byte("STOP\n"))
	}
}
