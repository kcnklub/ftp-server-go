package ftp

import (
	"fmt"
	"net"
	"os"
)

func (c *FTPConn) retr(args []string) {

	filename := args[0]

	file, err := os.ReadFile(fmt.Sprintf(c.Root + "/" + filename))
	if err != nil {
		c.respond(status426)
	}

	c.respond(status150)

	dataConn, err := net.Dial("tcp", c.DataAddr)
	if err != nil {
		c.respond("Cannot open DTP")
	}
	defer dataConn.Close()

	dataConn.Write(file)

	fmt.Fprint(dataConn, "\n")
	c.respond(status226)
}
