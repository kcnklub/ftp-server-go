package ftp

import (
	"bufio"
	"log"
	"net"
	"testing"
	"time"
)

/*
This test file tests FTP commands that use the DTP.
*/

func TestListCommand(t *testing.T) {
	expected := [4]string{status220, status200, status150, status226}
	server, client := net.Pipe()

	go func(t *testing.T) {
		go readAndAssertFromServer(t, client, expected[:])

		go func() {

			l, _ := net.Listen("tcp", ":8080")
			clientDtp, _ := l.Accept()

			scanner := bufio.NewReader(clientDtp)
			buf, _, err := scanner.ReadLine()
			if err != nil {
				log.Printf("TEST READLINE ERROR: %s\n", err)
			}

			output := string(buf[:])
			log.Println(output)
			if output != "test.txt" {
				t.Fatalf("Expected \"%s\" but got \"%s\"", "test.txt", output)
			}

		}()

		client.Write([]byte("PORT 127,0,0,1,31,144\n"))
		client.Write([]byte("LIST\n"))
		time.Sleep(2 * time.Second)
		client.Close()
	}(t)

	ftpConn := NewConn(server)
	ftpConn.Root = "./../test_content"
	Serve(&ftpConn)
}

func TestListNoDataConnectionCommand(t *testing.T) {
	expected := [4]string{status220, status150, status426}
	server, client := net.Pipe()

	go func(t *testing.T) {
		go readAndAssertFromServer(t, client, expected[:])

		client.Write([]byte("LIST\n"))
		time.Sleep(2 * time.Second)
		client.Close()
	}(t)

	ftpConn := NewConn(server)
	ftpConn.Root = "./../test_content"
	Serve(&ftpConn)
}
