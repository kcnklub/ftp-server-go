package ftp

import (
	"bufio"
	"log"
	"net"
	"os"
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
			defer l.Close()
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
		time.Sleep(100 * time.Millisecond)
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
		time.Sleep(100 * time.Millisecond)
		client.Close()
	}(t)

	ftpConn := NewConn(server)
	ftpConn.Root = "./../test_content"
	Serve(&ftpConn)
}

func TestListInvalidDirectorySelected(t *testing.T) {
	expected := [4]string{status220, status550}
	server, client := net.Pipe()

	go func(t *testing.T) {
		go readAndAssertFromServer(t, client, expected[:])

		client.Write([]byte("LIST\n"))
		time.Sleep(100 * time.Millisecond)
		client.Close()
	}(t)

	ftpConn := NewConn(server)
	ftpConn.Root = "./../test_content/I_DONT_EXIST"
	Serve(&ftpConn)
}

func TestRetrCommand(t *testing.T) {
	expected := [4]string{status220, status200, status150, status226}
	server, client := net.Pipe()

	go func(t *testing.T) {
		go readAndAssertFromServer(t, client, expected[:])

		go func() {
			l, _ := net.Listen("tcp", ":8080")
			defer l.Close()
			clientDtp, _ := l.Accept()

			scanner := bufio.NewReader(clientDtp)
			buf, _, err := scanner.ReadLine()
			if err != nil {
				log.Printf("TEST READLINE ERROR: %s\n", err)
			}

			output := string(buf[:])
			log.Println(output)
			if output != "test content" {
				t.Fatalf("Expected \"%s\" but got \"%s\"", "test content", output)
			}
		}()

		client.Write([]byte("PORT 127,0,0,1,31,144\n"))
		client.Write([]byte("RETR test.txt\n"))
		time.Sleep(100 * time.Millisecond)
		client.Close()
	}(t)

	ftpConn := NewConn(server)
	ftpConn.Root = "./../test_content"
	Serve(&ftpConn)

}

func TestStorCommand(t *testing.T) {
	expected := [4]string{status220, status200, status150, status226}
	server, client := net.Pipe()

	go func(t *testing.T) {
		go readAndAssertFromServer(t, client, expected[:])

		go func() {
			l, _ := net.Listen("tcp", ":8080")
			defer l.Close()
			clientDtp, _ := l.Accept()
			defer clientDtp.Close()

			clientDtp.Write([]byte("This is in the new file"))
		}()

		client.Write([]byte("PORT 127,0,0,1,31,144\n"))
		client.Write([]byte("STOR uploaded_test.txt\n"))
		time.Sleep(100 * time.Millisecond)
		client.Close()
	}(t)

	ftpConn := NewConn(server)
	ftpConn.Root = "./../test_content"
	Serve(&ftpConn)

	file, _ := os.ReadFile("./../test_content/uploaded_test.txt")
	output := string(file)
	if output != "This is in the new file" {
		t.Fatalf("Expected \"%s\" but got \"%s\"", "This is in this new file", output)
	}

	if err := os.Remove("./../test_content/uploaded_test.txt"); err != nil {
		t.Fatalf("Failed cleaning up test contents file: %s\n", err)
	}
}
