package ftp

import (
	"net"
	"testing"
)

func TestList(t *testing.T) {

	//output := NewConn(r)

	l, _ := net.Listen("tcp", ":8081")
	go func() {

		client, _ := net.Dial("tcp", ":8081")
		defer client.Close()
		client.Write([]byte("USER testuser"))
	}()

	c, _ := l.Accept()

	ftpConn := NewConn(c)

	if ftpConn.root != "." {
		t.Fatalf("failed to initialize ftpConn")
	}

	Serve(ftpConn)
}
