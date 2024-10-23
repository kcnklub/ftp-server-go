package ftp

import (
	"io"
)

type FTPConn struct {
	Conn     io.ReadWriteCloser
	DataAddr string
	dataType string
	root     string
}

func NewConn(conn io.ReadWriteCloser) FTPConn {
	return FTPConn{conn, "", "ascii", "."}
}
