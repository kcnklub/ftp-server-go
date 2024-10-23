package ftp

import (
	"fmt"
	"net"
	"os"
)

func (c *FTPConn) list() {
	dir, err := os.ReadDir(c.root)
	if err != nil {
		c.respond("Error reading directory")
	}

	c.respond(status150)

	dataConn, err := net.Dial("tcp", c.DataAddr)
	if err != nil {
		c.respond("Cannot open DTP")
	}
	defer dataConn.Close()

	for _, entry := range dir {
		fmt.Fprint(dataConn, entry.Name(), "\n")
	}

	fmt.Fprint(dataConn, "\n")
	c.respond(status226)
}
