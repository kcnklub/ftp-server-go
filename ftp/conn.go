package ftp

import "net"

type FTPConn struct {
	Conn     net.Conn
	DataConn net.Conn
	dataType string
}

func NewConn(conn net.Conn) FTPConn {
	return FTPConn{conn, nil, "ascii"}
}
