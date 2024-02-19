package utils

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

// FatalApplication global error application
func FatalApplication(msg string, err error) {
	log.Fatalf("%s > %s\n", msg, err)
}

// ReadMessage reads a message from the connection.
func ReadMessage(conn net.Conn) (msg []byte, err error) {
	var length uint64
	if err = binary.Read(conn, binary.BigEndian, &length); err != nil {
		return
	}

	msg = make([]byte, length)
	_, err = io.ReadFull(conn, msg)
	return
}

// WriteMessage writes a message to the connection.
func WriteMessage(conn net.Conn, msg []byte) (err error) {
	if err = binary.Write(conn, binary.BigEndian, uint64(len(msg))); err != nil {
		return
	}
	_, err = conn.Write(msg)
	return
}

func WriteInt64(conn net.Conn, num int64) (err error) {
	return binary.Write(conn, binary.LittleEndian, num)
}

func ReadInt64(conn net.Conn) (int64, error) {
	var num int64
	err := binary.Read(conn, binary.LittleEndian, &num)
	if err != nil {
		return 0, err
	}

	return num, nil
}
