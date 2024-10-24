package ftp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func (c *FTPConn) stor(args []string) {

	filename := args[0]

	f, err := os.Create(fmt.Sprintf(c.Root + "/" + filename))
	if err != nil {
		c.respond(status426)
		return
	}
	defer f.Close()

	c.respond(status150)

	d, err := net.Dial("tcp", c.DataAddr)
	if err != nil {
		c.respond("Cannot open DTP")
		return
	}
	defer d.Close()

	buf := make([]byte, 512)
	r := bufio.NewReader(d)
	i, err := r.Read(buf)
	if err != nil {
		c.respond(status426)
		return
	}

	log.Printf("Reading %d bytes\n", i)
	log.Printf("Reading %s\n", string(buf))

	i, err = f.Write(buf[:i])
	if err != nil {
		c.respond(status426)
		return
	}
	log.Printf("Writing %d bytes\n", i)

	c.respond(status226)
}
