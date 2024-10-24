package main

import (
	"example.com/ftp-server/ftp"
	"fmt"
	"net"
	"os"
)

func main() {

	ln, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open listener")
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to open connection")
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()
			c := ftp.NewConn(conn)
			ftp.Serve(&c)
		}(conn)
	}
}
