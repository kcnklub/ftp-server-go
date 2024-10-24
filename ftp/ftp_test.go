package ftp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"testing"
	"time"
)

/*
This file tests all non DTP commands in the FTP protocol
*/

func TestUserCommand(t *testing.T) {
	expected := [2]string{status220, fmt.Sprintf(status230, "testuser")}
	server, client := net.Pipe()

	go func(t *testing.T) {
		go readAndAssertFromServer(t, client, expected[:])

		client.Write([]byte("USER testuser\n"))
		time.Sleep(500 * time.Millisecond)
		client.Close()
	}(t)

	ftpConn := NewConn(server)
	Serve(&ftpConn)
}

func TestQuitCommand(t *testing.T) {
	expected := [2]string{status220, status221}
	server, client := net.Pipe()

	go func(t *testing.T) {
		go readAndAssertFromServer(t, client, expected[:])

		client.Write([]byte("QUIT\n"))
		time.Sleep(500 * time.Millisecond)
		client.Close()
	}(t)

	ftpConn := NewConn(server)
	Serve(&ftpConn)
}

func TestNotSupportedCommand(t *testing.T) {
	expected := [2]string{status220, status502}
	server, client := net.Pipe()

	go func(t *testing.T) {
		go readAndAssertFromServer(t, client, expected[:])

		client.Write([]byte("NONSUPPORTED\n"))
		time.Sleep(500 * time.Millisecond)
		client.Close()
	}(t)

	ftpConn := NewConn(server)
	Serve(&ftpConn)
}
func TestPortCommand(t *testing.T) {
	expected := [2]string{status220, status200}
	server, client := net.Pipe()

	go func(t *testing.T) {
		go readAndAssertFromServer(t, client, expected[:])

		client.Write([]byte("PORT 127,0,0,1,31,144\n"))
		time.Sleep(500 * time.Millisecond)
		client.Close()
	}(t)

	ftpConn := NewConn(server)
	Serve(&ftpConn)

	if ftpConn.DataAddr != "127.0.0.1:8080" {
		t.Fatalf("Failed to set port: expected: %s, got: %s", "127.0.0.1:8080", ftpConn.DataAddr)
	}
}

func readAndAssertFromServer(t *testing.T, client net.Conn, expected []string) {
	scanner := bufio.NewReader(client)
	count := 0
	for {

		buf, _, err := scanner.ReadLine()
		if err != nil {
			log.Printf("TEST READLINE ERROR: %s\n", err)
			break
		}

		output := string(buf[:])
		log.Println(output)
		if output != expected[count] {
			t.Fatalf("Expected \"%s\" for the %d message but got \"%s\"", expected[count], count, output)
		}
		count++
	}
}
