package ftp

import "net"

type FTPConn struct {
	Conn     net.Conn
	DataAddr string
	dataType string
	root     string
}

func NewConn(conn net.Conn) FTPConn {
	return FTPConn{conn, "", "ascii", "."}
}
